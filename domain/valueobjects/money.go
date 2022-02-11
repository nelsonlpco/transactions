package valueobjects

import (
	"errors"
	"math"
)

var ErrorInvalidAmmount = errors.New(`"the amount must be different of zero"`)

type Money float64

func NewMoney(amount float64) Money {
	return Money(amount)
}

func (m *Money) Validate() error {
	if *(m) == 0 {
		return ErrorInvalidAmmount
	}

	return nil
}

func (m *Money) Format(precision int) {
	value := float64(*m)
	precisionFactor := float64(math.Pow(10, float64(precision)))
	value = math.Round(value*precisionFactor) / precisionFactor

	*m = Money(value)
}

func (m *Money) ToFloat64() float64 {
	return float64(*m)
}
