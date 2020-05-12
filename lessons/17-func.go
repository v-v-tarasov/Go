//функции - переменное число аргументов функции

package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		fmt.Println(total, v)
		total += v
	}
	return total
}

func main() {
	fmt.Println(add(1, 2, 3, 4, 5))
}
