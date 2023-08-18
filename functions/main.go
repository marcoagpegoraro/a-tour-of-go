package main

import "fmt"

// When two or more consecutive named function parameters
// share a type, you can omit the type from all but the last.
// func add(x, y int) int {
func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
// These names should be used to document the meaning of the return values.
// A return statement without arguments returns the named return values. This is known as a "naked" return.
// Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// fmt.Println(add(10, 23))
	// fmt.Println(swap("Olá", "Mundo"))
	fmt.Println(split(17))
}
