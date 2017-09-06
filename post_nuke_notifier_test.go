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
)

func TestSendPostNukeNotifications(t *testing.T) {

	fakeUserEmail := "fakeser@fake.co"

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
				},
				DeleteAccessKeyOutput: &iam.DeleteAccessKeyOutput {},
			},
		},
	}

	// Create mock mailgun and configure it to use that
	mockMailgun := NewMockMailGun()

	// Call post nuke notifier
	result, err := SendPostNukeMockNotifications(mockMailgun, params)
	assert.NoError(t, err, fmt.Sprintf("Unexpected error"))

	// Wait until mock mailgun is invoked
	msg, err := mockMailgun.WaitForNextMessage(time.Second)
	assert.NoError(t, err, fmt.Sprintf("Unexpected error"))
	log.Printf("Mailgun message: %v", msg)

	// Verify that the correct outputs are returned (should propagate checkpoints)
	log.Printf("result: %v", result)
	assert.EqualValues(t, len(result.NukedKeyEvents), len(params.NukedKeyEvents))
	assert.EqualValues(t, result.NukedKeyEvents, params.NukedKeyEvents)


}
