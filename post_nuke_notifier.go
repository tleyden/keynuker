// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mailgun/mailgun-go.v1"
)

type ParamsPostNukeNotifier struct {

	// These fields are inputs from the upstream nuke-leaked-aws-keys action
	Id                     string `json:"_id"`
	NukedKeyEvents         []NukedKeyEvent
	GithubEventCheckpoints GithubEventCheckpoints

	// The FROM address that will be used for any notifications
	EmailFromAddress string

	// Optionally specify the Keynuker admin email to be CC'd about any leaked/nuked keys
	KeynukerAdminEmailCCAddress string

}

func (p ParamsPostNukeNotifier) Validate() error {
	if p.EmailFromAddress == "" {
		return fmt.Errorf("You must specify an EmailFromAddress")
	}
	if p.KeynukerAdminEmailCCAddress == "" {
		return fmt.Errorf("You must specify a KeynukerAdminEmailCCAddress")
	}
	return nil
}

type ResultPostNukeNotifier struct {
	Id                     string `json:"_id"`
	NukedKeyEvents         []NukedKeyEvent
	GithubEventCheckpoints GithubEventCheckpoints

	// Mailgun delivery id's for messages
	DeliveryIds []string
}

// Entry point with dependency injection that takes a mailer object, might be live mailgun endpoint or mock
func SendPostNukeNotifications(mailer mailgun.Mailgun, params ParamsPostNukeNotifier) (result ResultPostNukeNotifier, err error) {

	if err := params.Validate(); err != nil {
		return result, err
	}

	// Propagate params -> result
	result.NukedKeyEvents = params.NukedKeyEvents
	result.GithubEventCheckpoints = params.GithubEventCheckpoints
	result.Id = params.Id

	for _, nukedKeyEvent := range params.NukedKeyEvents {

		if nukedKeyEvent.LeakedKeyEvent.GithubEvent == nil ||
			nukedKeyEvent.DeleteAccessKeyOutput == nil ||
			nukedKeyEvent.LeakedKeyEvent.AccessKeyMetadata.AccessKeyId == nil {
			log.Printf("Warning: invalid nil params.  Skipping notification for nukedKeyEvent: %+v", nukedKeyEvent)
			continue
		}

		// If the GithubUser.Email field is not available, fallback to admin user
		recipient := params.KeynukerAdminEmailCCAddress

		// Use github user email if present
		if nukedKeyEvent.LeakedKeyEvent.LeakerEmail != "" {
			recipient = nukedKeyEvent.LeakedKeyEvent.LeakerEmail
		} else {
			log.Printf(
				"Could not discover email for github user: %+v.  Falling back to: %v.  LeakedKeyEvent: %+v",
				nukedKeyEvent.LeakedKeyEvent.GithubUser,
				recipient,
				nukedKeyEvent.LeakedKeyEvent,
			)
		}

		messageBody := fmt.Sprintf(
			"Dear %v, looks like on %v you leaked an AWS key: %+v via github event: %+v. "+
				"The AWS key was attempted to be deleted on AWS, with AWS returning: %+v.",
			recipient,
			nukedKeyEvent.NukedOn,
			*nukedKeyEvent.LeakedKeyEvent.AccessKeyMetadata.AccessKeyId,
			*nukedKeyEvent.LeakedKeyEvent.GithubEvent,
			*nukedKeyEvent.DeleteAccessKeyOutput,
		)

		log.Printf("Message body: %v.  Recipient: %v.  From: %v", messageBody, recipient, params.EmailFromAddress)

		message := mailgun.NewMessage(
			params.EmailFromAddress,
			"WARNING: An AWS key was leaked from your Github account.  Detected and nuked! 🔐💥",
			messageBody,
			recipient,
		)

		if recipient != params.KeynukerAdminEmailCCAddress && params.KeynukerAdminEmailCCAddress != "" {
			message.AddCC(params.KeynukerAdminEmailCCAddress)
		}

		mes, id, err := mailer.Send(message)
		if err != nil {
			return result, fmt.Errorf("Error sending message: %v.  Mes: %v", err, mes)
		}

		result.DeliveryIds = append(result.DeliveryIds, id)

		log.Printf("Delivery id: %v for outgoing email message: %+v", id, message)

	}

	return result, nil

}

// Entry point when using actual OpenWhisk action.  Uses live mailgun endpoint.
func SendPostNukeMailgunNotifications(params ParamsPostNukeNotifier) (result ResultPostNukeNotifier, err error) {

	// TODO: create mailgun endpoint using credentials found in parameters

	return SendPostNukeNotifications(nil, params)
}

// Entry point when using test.  Uses mock mailgun endpoint.
func SendPostNukeMockNotifications(mockMailgun mailgun.Mailgun, params ParamsPostNukeNotifier) (result ResultPostNukeNotifier, err error) {


	return SendPostNukeNotifications(mockMailgun, params)

}

func NewMailgunFromEnvironmentVariables() (mg mailgun.Mailgun, err error) {

	mailerDomain := os.Getenv("MAILERDOMAIN")
	if mailerDomain == "" {
		return nil, fmt.Errorf("You must set MAILERDOMAIN env variable")
	}
	mailerAPIKey := os.Getenv("MAILERAPIKEY")
	if mailerAPIKey == "" {
		return nil, fmt.Errorf("You must set MAILERAPIKEY env variable")
	}
	mailerPublicAPIKey := os.Getenv("MAILERPUBLICAPIKEY")
	if mailerPublicAPIKey == "" {
		return nil, fmt.Errorf("You must set MAILERPUBLICAPIKEY env variable")
	}

	return mailgun.NewMailgun(
		mailerDomain,
		mailerAPIKey,
		mailerPublicAPIKey,
	), nil

}
