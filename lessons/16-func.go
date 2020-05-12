// функции

package main

import "fmt"

var xs string = "return variable"

func main() {
	fmt.Println("run func main")
	fmt.Println(f1())
}

func f1() string {
	fmt.Println("run func f1")
	return f2()
}

func f2() string {
	fmt.Println("run func f2")
	return xs
}
