package main

import (
	"fmt"
	"runtime"
	"time"
)

func greetings() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	greetings()

	os := runtime.GOOS

	switch os {
	case "darwin":
		fmt.Println("runnning on macOS")
	case "linux":
		fmt.Println("Running on linux")
	default:
		fmt.Println(fmt.Sprintf("%s", os))
	}
}
