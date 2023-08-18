package main

import "fmt"

func main() {
	sum := 0

	//FOR LOOP
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }

	//For continued (while in C/Java)
	sum = 1
	for sum < 1000 {
		sum += sum
	}

	// for true {
	for {
		fmt.Println("Loop infinito, ao n ser q vc use 'break'")
		break
	}
	fmt.Println(sum)
}
