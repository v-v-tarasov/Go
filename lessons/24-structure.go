// Структуры и указатели на структуры

package main

import "fmt"

type person struct{
    name string
    age int
}

func main() {
    var tom = person {name: "Tom", age: 18}
    fmt.Println("Name:", tom.name)
    fmt.Println("Age:", tom.age)

    tom.age = 50
    fmt.Println(tom.name, tom.age)

    fmt.Println("Pointers to structures:")

    var tomPointer *person = &tom
    tomPointer.age = 25
    fmt.Println("New age:", tom.age)
    (*tomPointer).age = 45
    fmt.Println("New-New age:", tom.age)
}
