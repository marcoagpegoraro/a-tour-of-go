package main

import "fmt"

func main() {
	height := 1.79

	if height < 2.00 {
		fmt.Println("Your size is normal")
	} else if height < 1.50 {
		fmt.Println("You are small")
	} else {
		fmt.Println("Your size is medium")
	}

}
