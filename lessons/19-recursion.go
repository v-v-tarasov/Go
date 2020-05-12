// рекурсия

package main

import "fmt"

func main() {
	fmt.Println(factorial(1))
	fmt.Println("run main")
}

func factorial(x uint) uint {
	fmt.Println("run factorial")
	if x == 0 {
		fmt.Println("run if")
		return 1
	}
	fmt.Println("run return factorial")
	return x * factorial(x-1)
}
