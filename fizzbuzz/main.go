package main

import "fmt"

func main() {
	for a := 1; a <= 30; a++ {
		if a % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if a % 3 == 0 {
			fmt.Println("Fizz")
		} else if a % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(a)
		}
	}
}