// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"fmt"
	"log"
	"os"

	"encoding/json"
	"github.com/tleyden/keynuker/keynuker-go-common"
	"gopkg.in/mailgun/mailgun-go.v1"
)

// Mailer (Mailgun) Params
type MailerParams struct {

	// The Mailgun API key for notifications
	ApiKey string `json:"mailer_api_key"`

	// The Mailgun public api key.
	PublicApiKey string `json:"mailer_public_api_key"`

	// The mailgun domain
	Domain string `json:"mailer_domain"`
}

type ParamsPostNukeNotifier struct {

	// MailerParams
	MailerParams

	// This is the name of the KeyNuker "org/tenant".  Defaults to "default", but allows to be extended multi-tenant.
	KeyNukerOrg string

	// These fields are inputs from the upstream nuke-leaked-aws-keys action
	NukedKeyEvents         []NukedKeyEvent
	GithubEventCheckpoints GithubEventCheckpoints

	// The FROM address that will be used for any notifications
	EmailFromAddress string `json:"email_from_address"`

	// Optionally specify the Keynuker admin email to be CC'd about any leaked/nuked keys
	KeynukerAdminEmailCCAddress string `json:"admin_email_cc_address"`
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
	NukedKeyEvents         []NukedKeyEvent
	GithubEventCheckpoints GithubEventCheckpoints

	// Mailgun delivery id's for messages
	DeliveryIds []string
}

type DocumentWrapperPostNukeNotifier struct {
	// Serialize into a form that the cloudant db adapter expects
	Doc   ResultPostNukeNotifier `json:"doc"`
	DocId string                 `json:"docid"`
}

// Entry point with dependency injection that takes a mailer object, might be live mailgun endpoint or mock
func SendPostNukeNotifications(mailer mailgun.Mailgun, params ParamsPostNukeNotifier) (result ResultPostNukeNotifier, err error) {

	if err := params.Validate(); err != nil {
		return result, err
	}

	// Propagate params -> result
	result.NukedKeyEvents = params.NukedKeyEvents
	result.GithubEventCheckpoints = params.GithubEventCheckpoints

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

		githubEventJson, _ := json.MarshalIndent(nukedKeyEvent.LeakedKeyEvent.GithubEvent, "", "    ")
		deleteAccessKeyJson, _ := json.MarshalIndent(nukedKeyEvent.DeleteAccessKeyOutput, "", "    ")

		messageBody := fmt.Sprintf(
			"Dear %v, looks like you leaked an AWS key: %+v via github event: %+v. "+
				"The AWS key was attempted to be deleted on AWS, with AWS result: %+v.  Timestamp: %v",
			recipient,
			*nukedKeyEvent.LeakedKeyEvent.AccessKeyMetadata.AccessKeyId,
			string(githubEventJson),
			string(deleteAccessKeyJson),
			nukedKeyEvent.NukedOn,
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

	mailer := mailgun.NewMailgun(
		params.MailerParams.Domain,
		params.MailerParams.ApiKey,
		params.MailerParams.PublicApiKey,
	)

	return SendPostNukeNotifications(mailer, params)
}

// Entry point when using test.  Uses mock mailgun endpoint.
func SendPostNukeMockNotifications(mockMailgun mailgun.Mailgun, params ParamsPostNukeNotifier) (result ResultPostNukeNotifier, err error) {

	return SendPostNukeNotifications(mockMailgun, params)

}

func NewMailgunFromEnvironmentVariables() (mg mailgun.Mailgun, err error) {

	mailerDomain := os.Getenv(keynuker_go_common.EnvVarKeyNukerMailerDomain)
	if mailerDomain == "" {
		return nil, fmt.Errorf("You must set %v env variable", keynuker_go_common.EnvVarKeyNukerMailerDomain)
	}
	mailerAPIKey := os.Getenv(keynuker_go_common.EnvVarKeyNukerMailerApiKey)
	if mailerAPIKey == "" {
		return nil, fmt.Errorf("You must set %v env variable", keynuker_go_common.EnvVarKeyNukerMailerApiKey)
	}
	mailerPublicAPIKey := os.Getenv(keynuker_go_common.EnvVarKeyNukerMailerPublicApiKey)
	if mailerPublicAPIKey == "" {
		return nil, fmt.Errorf("You must set %v env variable", keynuker_go_common.EnvVarKeyNukerMailerPublicApiKey)
	}

	return mailgun.NewMailgun(
		mailerDomain,
		mailerAPIKey,
		mailerPublicAPIKey,
	), nil

}
