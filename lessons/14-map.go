// карты

package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	num := make(map[string]string)
	num["0"] = "0"
	num["1"] = "1"
	num["2"] = "2"
	num["3"] = "3"
	num["4"] = "4"
	fmt.Println(num["0"], num["2"], num["4"])

	num2 := map[string]string{
		"0": "0",
		"1": "1",
		"2": "2",
		"3": "3",
		"4": "4",
	}
	fmt.Println(num2["1"], num2["3"], num["0"])
}
