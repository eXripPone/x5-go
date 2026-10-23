package main

import "fmt"

func main() {
	var n int
	total := 0
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Некорректный ввод")
		return
	}
	mas := make([]int, n)	
	for i := range mas {
		fmt.Scan(&mas[i])
		total += mas[i]
	}
	min, max := mas[0], mas[0]
	for _, i := range mas {
	if i < min {
		min = i
	}
	if i > max {
		max = i
	}
	}
	fmt.Println("минимум:", min, "максимум:", max, "сумма:", total, "среднее:", float64(total)/float64(n))
}