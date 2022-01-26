package mathematics

import (
	"errors"
	"math"
)

type Mathematics struct{}

func (m *Mathematics) Addition(valueA, valueB float64) (result float64, err error) {
	result = valueA + valueB
	err = getInfiniteError(result)
	return
}

func (m *Mathematics) Subtraction(valueA, valueB float64) (result float64, err error) {
	if !checkSignal(valueB) {
		result, err = m.Addition(valueA, -valueB)
		return
	}

	result = valueA - valueB
	err = getInfiniteError(result)
	return
}

func (m *Mathematics) Multiplication(valueA, valueB float64) (result float64, err error) {
	result = valueA * valueB
	err = getInfiniteError(result)
	return
}

func (m *Mathematics) Division(valueA, valueB float64) (result float64, err error) {
	if valueB == 0.0 {
		err = errors.New("division by zero")
		return
	}

	result = valueA / valueB
	err = getInfiniteError(result)
	return
}

func checkSignal(value float64) bool {
	return !math.Signbit(value) // returns false for positive
}

func getInfiniteError(result float64) (err error) {
	if math.IsInf(result, 0) {
		err = errors.New("result overloads type limit")
		return
	}
	return
}
