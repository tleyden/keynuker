package keynuker

import (
	"log"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/aws/aws-sdk-go/service/iam"
	"gopkg.in/mailgun/mailgun-go.v1"
	"encoding/json"
)


func TestSendPostNukeNotifications(t *testing.T) {

	fakeUserEmail := "fakeUser@fake.co"
	fakeAdminEmail := "fakeAdmin@fake.co"

	// Mailgun instance -- either live or mock depending on whether integration tests enabled
	var mg mailgun.Mailgun
	var err error

	// Create fake inputs that include nuked key events
	params := ParamsPostNukeNotifier{
		NukedKeyEvents: []NukedKeyEvent{
			{
				NukedOn: time.Now(),
				LeakedKeyEvent: LeakedKeyEvent{
					GithubUser: &github.User{
						Email: aws.String(fakeUserEmail),
					},
					GithubEvent: &github.Event{},
					AccessKeyMetadata: FetchedAwsAccessKey{
						AccessKeyId: aws.String("fake-aws-access-key"),
					},
				},
				DeleteAccessKeyOutput: &iam.DeleteAccessKeyOutput {},
			},
		},
		KeynukerAdminEmailCCAddress: fakeAdminEmail,
		EmailFromAddress: fakeAdminEmail,
		MailerParams: MailerParams{},
	}

	// Unless integration tests are enabled, use mock
	mockTest := !IntegrationTestsEnabled()

	if mockTest {
		// Create mock mailgun and configure it to use that
		mg = NewMockMailGun()
	} else {

		// Integration test instructions -- set env variables MAILERDOMAIN, MAILERAPIKEY, and MAILERPUBLICAPIKEY

		mg, err = NewMailgunFromEnvironmentVariables()
		assert.NoError(t, err, fmt.Sprintf("Unexpected error"))

		fakeUserEmail = "youremail@your.org"  // Replace these to receive emails
		fakeAdminEmail = "fakeAdmin@your.org"

	}


	// Call post nuke notifier
	result, err := SendPostNukeMockNotifications(mg, params)
	assert.NoError(t, err, fmt.Sprintf("Unexpected error"))

	if mockTest {

		mockMailgun := mg.(*MockMailGun)

		// Wait until mock mailgun is invoked
		msg, err := mockMailgun.WaitForNextMessage(time.Second)
		assert.NoError(t, err, fmt.Sprintf("Unexpected error"))
		log.Printf("Mailgun message: %v", msg)

		// Verify that the correct outputs are returned (should propagate checkpoints)
		log.Printf("result: %v", result)
		assert.EqualValues(t, len(result.NukedKeyEvents), len(params.NukedKeyEvents))
		assert.EqualValues(t, result.NukedKeyEvents, params.NukedKeyEvents)

		// Should have at least one delivery id
		assert.True(t, len(result.DeliveryIds) >= 1)

	}



}
