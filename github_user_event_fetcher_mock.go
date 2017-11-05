/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/goautomock
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package keynuker

import (
	"fmt"
	mock "github.com/stretchr/testify/mock"

	context "context"
	github "github.com/google/go-github/github"
)

// GithubUserEventFetcherMock mock
type GithubUserEventFetcherMock struct {
	mock.Mock
}

func NewGithubUserEventFetcherMock() *GithubUserEventFetcherMock {
	return &GithubUserEventFetcherMock{}
}

// FetchUserEvents mocked method
func (m *GithubUserEventFetcherMock) FetchUserEvents(p0 context.Context, p1 FetchUserEventsInput) ([]*github.Event, error) {

	ret := m.Called(p0, p1)

	var r0 []*github.Event
	switch res := ret.Get(0).(type) {
	case nil:
	case []*github.Event:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ScanDownstreamContent mocked method
func (m *GithubUserEventFetcherMock) ScanDownstreamContent(p0 context.Context, p1 *github.Event, p2 []FetchedAwsAccessKey) ([]FetchedAwsAccessKey, error) {

	ret := m.Called(p0, p1, p2)

	var r0 []FetchedAwsAccessKey
	switch res := ret.Get(0).(type) {
	case nil:
	case []FetchedAwsAccessKey:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}
