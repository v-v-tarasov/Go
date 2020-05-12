// пользовательский ввод

package main

import "fmt"

func main() {
	fmt.Println("Введите целое беззнаковое число: ")
	var input uint8
	fmt.Scanf("%d", &input)
	fmt.Println("Введено: ", input)
	output := input * 2
	fmt.Println(output)
}

