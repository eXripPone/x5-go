package main

import "fmt"

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

	for _, sender := range channels {
		if email, ok := sender.(EmailNotifier); ok {
			fmt.Println("адрес email-канала:", email.Address)
		}
	}
}
