package main

import (
	"errors"
	"strings"
)

type EmailNotifier struct {
	BaseNotifier
	Address string
}

func (sender EmailNotifier) Send(text string) error {

	if strings.TrimSpace(sender.Address) == "" {
		return errors.New("email-адрес не задан")
	}
	sender.deliver(sender.Address, text)
	return nil
}
