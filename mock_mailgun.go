package keynuker

import (
	"fmt"
	"gopkg.in/mailgun/mailgun-go.v1"
	"time"
)

type MockMailGun struct {
	SentMessages chan *mailgun.Message

	// Embed the Mailgun interface. If unimplemented methods are called, it will panic
	mailgun.Mailgun
}

func NewMockMailGun() *MockMailGun {

	mockMailGun := MockMailGun{
		SentMessages: make(chan *mailgun.Message, 100),
	}

	return &mockMailGun
}

func (mmg *MockMailGun) Send(m *mailgun.Message) (string, string, error) {
	defer func() {
		mmg.SentMessages <- m
	}()
	return "", "", nil
}

func (mmg *MockMailGun) WaitForNextMessage(timeout time.Duration) (msg *mailgun.Message, err error) {
	select {
	case msg := <-mmg.SentMessages:
		return msg, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("Timed out waiting for message")
	}
}
