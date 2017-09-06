// Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements;
// and to You under the Apache License, Version 2.0.  See LICENSE in project root for full license + copyright.

package keynuker

import (
	"fmt"
	"gopkg.in/mailgun/mailgun-go.v1"
	"log"
)

type ParamsPostNukeNotifier struct {
	Id                     string `json:"_id"`
	NukedKeyEvents         []NukedKeyEvent
	GithubEventCheckpoints GithubEventCheckpoints

	EmailFromAddress string
}

type ResultPostNukeNotifier struct {
	Id                     string `json:"_id"`
	NukedKeyEvents         []NukedKeyEvent
	GithubEventCheckpoints GithubEventCheckpoints
}

// Entry point with dependency injection that takes a mailer object, might be live mailgun endpoint or mock
func SendPostNukeNotifications(mailer mailgun.Mailgun, params ParamsPostNukeNotifier) (result ResultPostNukeNotifier, err error) {

	// Propagate params -> result
	result.NukedKeyEvents = params.NukedKeyEvents
	result.GithubEventCheckpoints = params.GithubEventCheckpoints
	result.Id = params.Id

	for _, nukedKeyEvent := range params.NukedKeyEvents {

		if nukedKeyEvent.LeakedKeyEvent.GithubUser == nil ||
			nukedKeyEvent.LeakedKeyEvent.GithubEvent == nil ||
			nukedKeyEvent.DeleteAccessKeyOutput == nil {
			return result, fmt.Errorf("Invalid params")
		}

		recipient := *nukedKeyEvent.LeakedKeyEvent.GithubUser.Email

		messageBody := fmt.Sprintf(
			"Dear %v, looks like on %v you leaked an AWS key: %+v via github event: %+v. "+
				"The AWS key was attempted to be deleted on AWS, with AWS returning: %+v.",
			recipient,
			nukedKeyEvent.NukedOn,
			nukedKeyEvent.LeakedKeyEvent.AccessKeyMetadata,
			*nukedKeyEvent.LeakedKeyEvent.GithubEvent,
			*nukedKeyEvent.DeleteAccessKeyOutput,
		)

		log.Printf("Message body: %v.  Recipient: %v", messageBody, recipient)

		message := mailgun.NewMessage(
			params.EmailFromAddress,
			"WARNING: An AWS key was leaked from your Github account.  Detected and nuked 🔐💥!",
			messageBody,
			recipient,
		)

		_, id, err := mailer.Send(message)
		if err != nil {
			return result, err
		}

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
