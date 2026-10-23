package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Name  string
	Price int
}

type Order struct {
	ID     int
	Status string
	Items  []Item
}

var StatusNames = map[string]string{
	"new":       "Новый",
	"paid":      "Оплачен",
	"cancelled": "Отменён",
}

func NewOrder(id int) (*Order, error) {
	if id <= 0 {
		return nil, errors.New("ID должен быть положительным")
	}

	return &Order{ID: id, Status: "new"}, nil
}

func AddItem(order *Order, item Item) error {
	if order == nil {
		return errors.New("заказ не задан")
	}
	if strings.TrimSpace(item.Name) == "" {
		return errors.New("имя позиции пустое")
	}
	if item.Price < 0 {
		return errors.New("цена не может быть отрицательной")
	}

	order.Items = append(order.Items, item)
	return nil
}

func Total(order *Order) int {
	if order == nil {
		return 0
	}

	sum := 0
	for _, item := range order.Items {
		sum += item.Price
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите ID заказа: ")
	if !scanner.Scan() {
		return
	}

	id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("ID должен быть целым числом")
		return
	}

	order, err := NewOrder(id)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	for {
		fmt.Print("Название товара (стоп = завершить): ")
		if !scanner.Scan() {
			break
		}

		name := strings.TrimSpace(scanner.Text())
		if name == "стоп" {
			break
		}

		fmt.Print("Цена в руб.: ")
		if !scanner.Scan() {
			break
		}

		price, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("Цена должна быть целым числом")
			continue
		}

		if err := AddItem(order, Item{Name: name, Price: price}); err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Println("Товар добавлен")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения:", err)
	}

	fmt.Printf("\nЗаказ №%d, статус: %s\n", order.ID, StatusNames[order.Status])
	for _, item := range order.Items {
		fmt.Printf("%s = %d руб.\n", item.Name, item.Price)
	}
	fmt.Printf("Итого: %d руб.\n", Total(order))
}