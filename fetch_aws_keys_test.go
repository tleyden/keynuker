package keynuker

import (
	"encoding/json"
	"log"
	"testing"

	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"
	"github.com/tleyden/keynuker/keynuker-go-common"
)

//go:generate goautomock -template=testify -pkg github.com/aws/aws-sdk-go/service/iam/iamiface IAMAPI

func TestFetchIAMUsers(t *testing.T) {

	mockIAMService := NewIAMAPIMock()

	// -------------------- Setup mock: 1st invocation ------------------------------

	listUsersInputFirstInvocation := &iam.ListUsersInput{
		MaxItems: aws.Int64(1000),
	}

	// Create mock output with IsTruncated = true, meaning our code should make another request
	// to get the rest of the output
	mockListUsersOutputFirstInvocation := &iam.ListUsersOutput{
		IsTruncated: aws.Bool(true),
		Users: []*iam.User{
			{
				UserId: aws.String("fakeuser"),
			},
		},
		Marker: aws.String("2"),
	}

	mockIAMService.On("ListUsers", listUsersInputFirstInvocation).Return(
		mockListUsersOutputFirstInvocation,
		nil,
	).Once()

	// -------------------- Setup mock: 2nd invocation ------------------------------

	listUsersInputSecondInvocation := &iam.ListUsersInput{
		MaxItems: aws.Int64(1000),
		Marker: aws.String("2"),
	}

	// Create mock output with IsTruncated = true, meaning our code should make another request
	// to get the rest of the output
	mockListUsersOutputSecondInvocation := &iam.ListUsersOutput{
		IsTruncated: aws.Bool(false),
		Users: []*iam.User{
			{
				UserId: aws.String("fakeuser2"),
			},
		},
	}

	mockIAMService.On("ListUsers", listUsersInputSecondInvocation).Return(
		mockListUsersOutputSecondInvocation,
		nil,
	).Once()


	// -------------------- Invoke API call ------------------------------
	users, err := FetchIAMUsers(mockIAMService)
	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, 2, len(users))

}

func TestFetchAwsKeysViaSTSAssumeRole(t *testing.T) {

	SkipIfIntegrationsTestsNotEnabled(t)

	targetAwsAccounts, err := GetTargetAwsAccountsFromEnv()
	if err != nil {
		t.Skip("Error getting target aws accounts from env: %v", err)
	}

	initiatingAwsAcctCredsRaw, ok := os.LookupEnv(keynuker_go_common.EnvVarKeyNukerInitiatingAwsAccountCreds)
	if !ok {
		t.Fatalf("You must define environment variable %s", keynuker_go_common.EnvVarKeyNukerInitiatingAwsAccountCreds)
	}

	initiatingAwsAcctCreds := AwsCredentials{}
	if err := json.Unmarshal([]byte(initiatingAwsAcctCredsRaw), &initiatingAwsAcctCreds); err != nil {
		t.Fatalf("Error unmarshalling creds.  Err: %v", err)
	}

	params := ParamsFetchAwsKeys{
		KeyNukerOrg:                    "default",
		TargetAwsAccounts:              targetAwsAccounts,
		InitiatingAwsAccountAssumeRole: initiatingAwsAcctCreds,
	}

	doc, err := FetchAwsKeys(params)

	assert.NoError(t, err, "Unexpected Error")

	marshalled, err := json.MarshalIndent(doc, "", "    ")

	log.Printf("FetchedAwsKeys: %v", string(marshalled))

}
