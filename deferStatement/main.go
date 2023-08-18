package main

import "fmt"

func helloWorld() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func countTo10() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
func main() {
	helloWorld()
	countTo10()
}
