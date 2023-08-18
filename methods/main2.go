package main

import "fmt"

type Dog struct {
	name string
	age  int
}

func (dog Dog) bark() {
	fmt.Println("O cachorro " + dog.name + " está latindo")
}

func main() {
	dog := Dog{"Carlos", 3}

	dog.bark()
}
