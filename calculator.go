package main

import "errors"

// ErrZeroDivision is returned when second input is 0
var ErrZeroDivision = errors.New("can't divide by zero")

type calculatorService struct {
	number float64
}

func (c *calculatorService) sum(a, b float64) float64 {
	return a + b + c.number
}

func (c *calculatorService) sub(a, b float64) float64 {
	return a - b + c.number
}

func (c *calculatorService) multiply(a, b float64) float64 {
	return a*b + c.number
}

func (c *calculatorService) division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrZeroDivision
	}

	return a/b + c.number, nil
}
