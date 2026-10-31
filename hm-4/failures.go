package main

import (
	"errors"
	"fmt"
)

var ErrEmptyMessage = errors.New("пустое сообщение")

type SendError struct {
	Channel string
	Reason  string
}

func (failure *SendError) Error() string {
	return fmt.Sprintf("канал %s: %s", failure.Channel, failure.Reason)
}

type BrokenNotifier struct{}

func (BrokenNotifier) Send(string) error {
	panic("драйвер канала повреждён")
}
