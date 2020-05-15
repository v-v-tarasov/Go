// Встраиваемые типы ?или композиция?

package main

import "fmt"

<<<<<<< HEAD
type Person struct {
    Name string
}

func (p *Person) Introduce() {
    fmt.Println("Hi, my name is", p.Name)
}

type Say struct {
    *Person
    Power int
=======
func main() {
	fmt.Println("Start program...")
>>>>>>> 05f3fd459d44e7dcc157fc400fb8bf762335e6f3
}

func main() {
    goku := &Say{
        Person: &Person{"Goku"},
        Power: 9001,
    }
    goku.Introduce()
}

