package main

import "fmt"

func Notify(channel Notifier, text string) error {
	failure := safeSend(channel, text)
	if failure == nil {
		return nil
	}
	return fmt.Errorf("отправить уведомление: %w", failure)
}

func Dispatch(channels []Notifier, text string) error {
	for position := range channels {
		failure := Notify(channels[position], text)
		if failure != nil {
			return fmt.Errorf("канал %d: %w", position+1, failure)
		}
	}
	return nil
}
