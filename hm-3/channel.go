package main

import "fmt"

type Notifier interface {
	Send(string) error
}

type BaseNotifier struct {
	Channel string
}

func (base BaseNotifier) Log(text string) {
	fmt.Println(base.Channel, text)
}

func (base BaseNotifier) deliver(recipient, text string) {
	base.Log(fmt.Sprintf("to %s: %s", recipient, text))
}
