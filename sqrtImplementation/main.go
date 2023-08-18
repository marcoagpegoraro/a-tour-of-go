package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	lastNum, currentNum := 0.0, 1.0

	for {
		lastNum = currentNum
		fmt.Println(currentNum)
		currentNum -= (currentNum*currentNum - x) / (2 * currentNum)

		if lastNum == currentNum {
			return currentNum
		}
	}
}

func main() {
	fmt.Println("Digite um número para encontrarmos a raiz quadrada dele")
	var sqrtToCalc float64
	fmt.Scanf("%f", &sqrtToCalc)

	fmt.Println(Sqrt(sqrtToCalc))
}
