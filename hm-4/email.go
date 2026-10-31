package main

import (
	"strings"
)

type EmailNotifier struct {
	BaseNotifier
	Address string
}

func (sender EmailNotifier) Send(text string) error {

	if strings.TrimSpace(text) == "" {
		return ErrEmptyMessage
	}

	if strings.TrimSpace(sender.Address) == "" {
		return &SendError{Channel: "email", Reason: "email-адрес не задан"}
	}
	sender.deliver(sender.Address, text)
	return nil
}
