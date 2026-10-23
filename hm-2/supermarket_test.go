package main

import "testing"

func TestNewOrder(t *testing.T) {
	tests := []struct {
		id      int
		wantErr bool
	}{
		{1, false},
		{0, true},
		{-1, true},
	}

	for _, tt := range tests {
		order, err := NewOrder(tt.id)

		if (err != nil) != tt.wantErr {
			t.Errorf("ID %d: ошибка = %v", tt.id, err)
		}
		if !tt.wantErr && order == nil {
			t.Error("заказ не создан")
		}
	}
}

func TestAddItem(t *testing.T) {
	tests := []struct {
		name    string
		order   *Order
		item    Item
		wantErr bool
	}{
		{"обычный товар", &Order{ID: 1}, Item{"Хлеб", 50}, false},
		{"бесплатный товар", &Order{ID: 1}, Item{"Подарок", 0}, false},
		{"nil заказ", nil, Item{"Хлеб", 50}, true},
		{"пустое имя", &Order{ID: 1}, Item{"", 50}, true},
		{"пробелы", &Order{ID: 1}, Item{"   ", 50}, true},
		{"отрицательная цена", &Order{ID: 1}, Item{"Хлеб", -1}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AddItem(tt.order, tt.item)

			if (err != nil) != tt.wantErr {
				t.Fatalf("ошибка = %v, ожидалась = %v", err, tt.wantErr)
			}

			if tt.wantErr {
				if tt.order != nil && len(tt.order.Items) != 0 {
					t.Error("заказ изменился после ошибки")
				}
				return
			}

			if len(tt.order.Items) != 1 || tt.order.Items[0] != tt.item {
				t.Error("товар не добавлен")
			}
		})
	}
}

func TestTotal(t *testing.T) {
	tests := []struct {
		name  string
		items []Item
		want  int
	}{
		{"пустой заказ", nil, 0},
		{"один товар", []Item{{"Хлеб", 50}}, 50},
		{"два товара", []Item{{"Хлеб", 50}, {"Молоко", 90}}, 140},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := &Order{ID: 1}

			for _, item := range tt.items {
				if err := AddItem(order, item); err != nil {
					t.Fatal(err)
				}
			}

			if got := Total(order); got != tt.want {
				t.Errorf("сумма = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}

func TestOrderCopy(t *testing.T) {
	original := Order{ID: 1, Status: "new"}
	copied := original
	copied.Status = "paid"

	if original.Status != "new" {
		t.Error("исходный заказ изменился")
	}
}

func TestMap(t *testing.T) {
	statuses := make(map[string]string)
	statuses["new"] = "Новый"

	if statuses["new"] != "Новый" {
		t.Error("статус не записан")
	}
}