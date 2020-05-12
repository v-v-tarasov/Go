// пример использования массивов:
// вычисление среднего арифметического

package main

import "fmt"

func main() {
	var x [5] float32
	x[0] = 85
	x[1] = 90
	x[2] = 92
	x[3] = 89
	x[4] = 96

	var total float32 = 0
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
		total += x[i]
	}
	fmt.Println("Итого:", total / float32(len(x)))

	y := [5]float32{
		50,
		55,
		60,
		65,
		70,
	}
	total = 0
	for _, value := range y {
		fmt.Println(value)
		total += value
	}
	fmt.Println("Итого:", total / float32(len(y)))
}

