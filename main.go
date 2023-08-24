package main

import (
	"fmt"
	"go-test-automation-github-action/calc"
)

func main() {
	var a, b int
	fmt.Scanln(&a, &b)

	fmt.Println(calc.Add(3, 4))
	fmt.Println(calc.Sub(3, 4))
	fmt.Println(calc.Mul(3, 4))
	fmt.Println(calc.Div(3, 4))
}
