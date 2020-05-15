// указатели

package main

import "fmt"

func zero(xPtr *int) {
    *xPtr = 0
}

func nonZero(y int) {
    y = 25
}

func main() {
    x := 5
    y := 45
    zero(&x)
    fmt.Println(x)
    nonZero(y)
    fmt.Println(y)
}
