package main

import (
	"errors"
	"fmt"
)

func safeSend(channel Notifier, text string) (failure error) {
	defer func() {
		recovered := recover()
		if recovered == nil {
			return
		}
		switch cause := recovered.(type) {
		case error:
			failure = fmt.Errorf("паника при отправке: %w", cause)
		default:
			failure = fmt.Errorf("паника при отправке: %v", cause)
		}
	}()
	if channel == nil {
		return errors.New("канал уведомлений не задан")
	}
	return channel.Send(text)
}
