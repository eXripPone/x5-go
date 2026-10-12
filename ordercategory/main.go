package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a < 0 {
		fmt.Println("Некорректный ввод")
		return
	}
	if a < 10_000 {
		fmt.Println("small")
	} else if a < 50_000 {
		fmt.Println("regular")
	} else {
		fmt.Println("large")
	}
}