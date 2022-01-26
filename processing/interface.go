package processing

type ICalculator interface {
	Calculate(sentence string) (result float64, err error)
}
