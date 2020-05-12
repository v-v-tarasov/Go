// defer - отложенный вызов

package main

import "fmt"

func first() {
	fmt.Println("run func first")
}

func second() {
	fmt.Println("run func second")
}

func main() {
	fmt.Println("run func main")
	defer first()
	second()
}
