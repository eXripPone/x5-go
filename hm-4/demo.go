package main

import (
	"errors"
	"fmt"
)

func main() {
	channels := []Notifier{
		EmailNotifier{
			BaseNotifier: BaseNotifier{Channel: "email"},
			Address:      "student@example.com",
		},
		PushNotifier{
			BaseNotifier: BaseNotifier{Channel: "push"},
			DeviceID:     "device-42",
		},
	}

	if err := Dispatch(channels, "Заказ готов"); err != nil {
		fmt.Println("ошибка:", err)
		return
	}

	err := Dispatch(channels, "")
	fmt.Println("ошибка:", err)
	fmt.Println("пустое сообщение:", errors.Is(err, ErrEmptyMessage))

	err = Dispatch([]Notifier{PushNotifier{}}, "Заказ готов")
	var sendErr *SendError
	if errors.As(err, &sendErr) {
		fmt.Printf("не сработал канал %q, причина: %s\n", sendErr.Channel, sendErr.Reason)
	}

	err = Dispatch([]Notifier{BrokenNotifier{}}, "Заказ готов")
	fmt.Println("результат:", err)

	if err := Dispatch(channels, "Работа продолжается"); err != nil {
		fmt.Println("ошибка:", err)
	}
}
