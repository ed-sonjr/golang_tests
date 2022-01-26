package processing

import (
	"errors"
	"home/edsjr/work/testes/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var calculator ICalculator

func Test_Test_WhenInvalidSentenceShouldReturnsError(t *testing.T) {
	mathematicsMock := mocks.IMathematics{}
	calculator = &Calculator{mathematics: &mathematicsMock}

	_, err := calculator.Calculate("abacate")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid sentence", err.Error())

	_, err = calculator.Calculate("2..+1")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid sentence", err.Error())

	_, err = calculator.Calculate("+1.0+2")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid sentence", err.Error())

	_, err = calculator.Calculate("1.0-2*")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid sentence", err.Error())

	_, err = calculator.Calculate("1++2")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid sentence", err.Error())

	_, err = calculator.Calculate(("1.0.0+2"))
	assert.NotNil(t, err)
	assert.Equal(t, "invalid sentence", err.Error())
}

func Test_CalculateHighPrecedenceOperationsSuccess(t *testing.T) {
	mathematicsMock := mocks.IMathematics{}
	mathematicsMock.On("Multiplication", mock.Anything, mock.Anything).Return(6.0, nil)
	mathematicsMock.On("Division", mock.Anything, mock.Anything).Return(1.5, nil)
	calculator = &Calculator{mathematics: &mathematicsMock}

	result, err := calculator.Calculate("2.0*3")
	assert.Nil(t, err)
	assert.Equal(t, 6.0, result)

	result, err = calculator.Calculate("3.0/2.")
	assert.Nil(t, err)
	assert.Equal(t, 1.5, result)

	mathematicsMock = mocks.IMathematics{}
	mathematicsMock.On("Multiplication", mock.Anything, mock.Anything).Return(6.0, nil)
	mathematicsMock.On("Division", mock.Anything, mock.Anything).Return(3.0, nil)
	calculator = &Calculator{mathematics: &mathematicsMock}

	result, err = calculator.Calculate("2*3.0/2.0")
	assert.Nil(t, err)
	assert.Equal(t, 3.0, result)
}

func Test_OtherOperationsSuccess(t *testing.T) {
	mathematicsMock := mocks.IMathematics{}
	mathematicsMock.On("Addition", 2.0, 3.0).Return(5.0, nil)
	mathematicsMock.On("Subtraction", 2.0, 3.0).Return(-1.0, nil)
	mathematicsMock.On("Subtraction", mock.Anything, 4.0).Return(1.0, nil)
	calculator = &Calculator{mathematics: &mathematicsMock}

	result, err := calculator.Calculate("2.0+3")
	assert.Nil(t, err)
	assert.Equal(t, 5.0, result)

	result, err = calculator.Calculate("2.0-3.")
	assert.Nil(t, err)
	assert.Equal(t, -1.0, result)

	result, err = calculator.Calculate("2.0+3-4.0")
	assert.Nil(t, err)
	assert.Equal(t, 1.0, result)
}

func Test_MixedOperationsSuccess(t *testing.T) {
	mathematicsMock := mocks.IMathematics{}
	mathematicsMock.On("Addition", 1.0, mock.Anything).Return(8.0, nil)
	mathematicsMock.On("Multiplication", 2.0, 3.5).Return(7.0, nil)
	mathematicsMock.On("Division", 3.0, 2.0).Return(1.5, nil)
	mathematicsMock.On("Multiplication", 1.5, 4.0).Return(6.0, nil)
	mathematicsMock.On("Subtraction", 2.0, 6.0).Return(-4.0, nil)
	mathematicsMock.On("Subtraction", -4.0, 3.0).Return(-7.0, nil)
	calculator = &Calculator{mathematics: &mathematicsMock}

	result, err := calculator.Calculate("1.0+2*3.5")
	assert.Nil(t, err)
	assert.Equal(t, 8.0, result)

	result, err = calculator.Calculate("2.0-3/2.0*4-3")
	assert.Nil(t, err)
	assert.Equal(t, -7.0, result)
}

func Test_CalculateHighPrecedenceOperationsError(t *testing.T) {
	mathematicsMock := mocks.IMathematics{}
	mathematicsMock.On("Multiplication", mock.Anything, mock.Anything).Return(0.0, errors.New("result overloads type limit"))
	mathematicsMock.On("Division", mock.Anything, mock.Anything).Return(0.0, errors.New("result overloads type limit"))
	mathematicsMock.On("Division", mock.Anything, 0.0).Return(0.0, errors.New("division by zero"))
	calculator = &Calculator{mathematics: &mathematicsMock}

	_, err := calculator.Calculate("2.0*3")
	assert.NotNil(t, err)

	_, err = calculator.Calculate("3.0/3")
	assert.NotNil(t, err)

	_, err = calculator.Calculate("2.0/0")
	assert.NotNil(t, err)
}

func Test_CalculateOtherOperationsError(t *testing.T) {
	mathematicsMock := mocks.IMathematics{}
	mathematicsMock.On("Addition", mock.Anything, mock.Anything).Return(0.0, errors.New("result overloads type limit"))
	mathematicsMock.On("Subtraction", mock.Anything, mock.Anything).Return(0.0, errors.New("result overloads type limit"))
	calculator = &Calculator{mathematics: &mathematicsMock}

	_, err := calculator.Calculate("2.0+3")
	assert.NotNil(t, err)

	_, err = calculator.Calculate("3.0-2")
	assert.NotNil(t, err)
}
