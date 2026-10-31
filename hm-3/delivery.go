package main

import "errors"

func Dispatch(channels []Notifier, text string) error {
	for _, channel := range channels {
		if channel == nil {
			return errors.New("канал уведомлений не задан")
		}
		failure := channel.Send(text)
		if failure != nil {
			return failure
		}
	}
	return nil
}
