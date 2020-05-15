// Встраиваемые типы ?или композиция?

package main

import "fmt"

type Person struct {
    Name string
}

func (p *Person) Introduce() {
    fmt.Println("Hi, my name is", p.Name)
}

type Say struct {
    *Person
    Power int
}

func main() {
    goku := &Say{
        Person: &Person{"Goku"},
        Power: 9001,
    }
    goku.Introduce()
}

