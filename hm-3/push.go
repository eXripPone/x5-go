package main

import (
	"errors"
	"strings"
)

type PushNotifier struct {
	BaseNotifier
	DeviceID string
}

func (sender PushNotifier) Send(text string) error {

	if strings.TrimSpace(sender.DeviceID) == "" {
		return errors.New("устройство не задано")
	}
	sender.deliver(sender.DeviceID, text)
	return nil
}
