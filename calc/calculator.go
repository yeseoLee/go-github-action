package calc

import "errors"

var ErrDivisionByZero = errors.New("divide by zero")

func Add(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Mul(a int, b int) int {
	return a * b
}

func Div(a int, b int) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return float64(a) / float64(b), nil
}
