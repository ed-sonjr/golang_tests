package mathematics

type IMathematics interface {
	Addition(valueA, valueB float64) (result float64, err error)
	Subtraction(valueA, valueB float64) (result float64, err error)
	Multiplication(valueA, valueB float64) (result float64, err error)
	Division(valueA, valueB float64) (result float64, err error)
}
