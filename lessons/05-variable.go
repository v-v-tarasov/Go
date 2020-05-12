// переменные

package main

import "fmt"

var (
	firstVar = 3
	secondVar = 6
)

func main() {
	fmt.Println(secondVar / firstVar)
	var x string = "Hello World"
	fmt.Println(x)
	x = "Hello"
	fmt.Println(x)
	x = x + "World"
	fmt.Println(x)
	x = "Hello"
	var y string = "World"
	fmt.Println(x == y)
	fmt.Println(x != y)
	y = "Hello"
	fmt.Println(x == y)
	fmt.Println(x != y)
}

