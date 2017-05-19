package main

// CalculatorService hold business logic methods
type CalculatorService interface {
	sum(float64, float64) float64
	sub(float64, float64) float64
	multiply(float64, float64) float64
	division(float64, float64) (float64, error)
}
