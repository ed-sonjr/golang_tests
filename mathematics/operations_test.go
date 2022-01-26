package mathematics

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mathematics *Mathematics

func TestMain(m *testing.M) {
	mathematics = &Mathematics{}
	m.Run()
}

func Test_AdditionSuccess(t *testing.T) {
	result, err := mathematics.Addition(2.0, 1.0)
	assert.Nil(t, err)
	assert.Equal(t, 3.0, result)

	result, err = mathematics.Addition(-2.0, 1.0)
	assert.Nil(t, err)
	assert.Equal(t, -1.0, result)

	result, err = mathematics.Addition(2.0, -1.0)
	assert.Nil(t, err)
	assert.Equal(t, 1.0, result)
}
func Test_WhenAdding2NumbersOverloadingTypeBoundaryShouldBeReturnAnError(t *testing.T) {
	_, err := mathematics.Addition(math.MaxFloat64, math.MaxFloat64)
	assert.NotNil(t, err)
	assert.Equal(t, "result overloads type limit", err.Error())

	_, err = mathematics.Addition(-math.MaxFloat64, -math.MaxFloat64)
	assert.NotNil(t, err)
	assert.Equal(t, "result overloads type limit", err.Error())
}

func Test_SubtractionSuccess(t *testing.T) {
	result, err := mathematics.Subtraction(2.0, 1.0)
	assert.Nil(t, err)
	assert.Equal(t, 1.0, result)

	result, err = mathematics.Subtraction(-2.0, 1.0)
	assert.Nil(t, err)
	assert.Equal(t, -3.0, result)

	result, err = mathematics.Subtraction(2.0, -1.0)
	assert.Nil(t, err)
	assert.Equal(t, 3.0, result)
}

func Test_WhenSubtracting2NumbersOverloadingTypeBoundaryShouldBeReturnAnError(t *testing.T) {
	_, err := mathematics.Subtraction(-math.MaxFloat64, math.MaxFloat64)
	assert.NotNil(t, err)
	assert.Equal(t, "result overloads type limit", err.Error())

	_, err = mathematics.Subtraction(math.MaxFloat64, -math.MaxFloat64)
	assert.NotNil(t, err)
	assert.Equal(t, "result overloads type limit", err.Error())
}

func Test_MultiplicationSuccess(t *testing.T) {
	result, err := mathematics.Multiplication(2.0, 1.0)
	assert.Nil(t, err)
	assert.Equal(t, 2.0, result)

	result, err = mathematics.Multiplication(-2.0, 1.0)
	assert.Nil(t, err)
	assert.Equal(t, -2.0, result)

	result, err = mathematics.Multiplication(2.0, -1.0)
	assert.Nil(t, err)
	assert.Equal(t, -2.0, result)
}

func Test_WhenMultiplying2NumbersOverloadingTypeBoundaryShouldBeReturnAnError(t *testing.T) {
	_, err := mathematics.Multiplication(math.MaxFloat64, 2.0)
	assert.NotNil(t, err)
	assert.Equal(t, "result overloads type limit", err.Error())

	_, err = mathematics.Multiplication(-math.MaxFloat64, -2.0)
	assert.NotNil(t, err)
	assert.Equal(t, "result overloads type limit", err.Error())

	_, err = mathematics.Multiplication(-math.MaxFloat64, 2.0)
	assert.NotNil(t, err)
	assert.Equal(t, "result overloads type limit", err.Error())
}

func Test_DivisionSuccess(t *testing.T) {
	result, err := mathematics.Division(1.0, 2.0)
	assert.Nil(t, err)
	assert.Equal(t, 0.5, result)
}

func Test_WhenDividingAnyNumberByZeroShouldBeReturnsAnError(t *testing.T) {
	_, err := mathematics.Division(1.0, 0.0)
	assert.NotNil(t, err)
	assert.Equal(t, "division by zero", err.Error())
}

func Test_WhenDividingSmallestNumberByAnyNumberShouldBeReturnsZero(t *testing.T) {
	result, err := mathematics.Division(math.SmallestNonzeroFloat64, 2.0)
	assert.Nil(t, err)
	assert.Equal(t, 0.0, result)

	result, err = mathematics.Division(math.SmallestNonzeroFloat64, math.MaxFloat64)
	assert.Nil(t, err)
	assert.Equal(t, 0.0, result)
}
