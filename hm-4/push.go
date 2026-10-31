package main

import (
	"strings"
)

type PushNotifier struct {
	BaseNotifier
	DeviceID string
}

func (sender PushNotifier) Send(text string) error {

	if strings.TrimSpace(text) == "" {
		return ErrEmptyMessage
	}

	if strings.TrimSpace(sender.DeviceID) == "" {
		return &SendError{Channel: "push", Reason: "не задан device id"}
	}
	sender.deliver(sender.DeviceID, text)
	return nil
}
