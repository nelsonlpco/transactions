package valueobjects_test

import (
	"errors"
	"testing"

	"github.com/nelsonlpco/transactions/domain/valueobjects"
	"github.com/stretchr/testify/require"
)

type MoneyTestObject struct {
	Expected  valueobjects.Money
	Value     float64
	Precision int
}

func Test_should_be_create_a_money(t *testing.T) {
	amount := valueobjects.NewMoney(10)

	require.NotNil(t, amount)
	require.Nil(t, amount.Validate())
}

func Test_should_be_get_money_to_float64(t *testing.T) {
	expectedValue := 10.43
	amount := valueobjects.Money(expectedValue)

	require.Equal(t, expectedValue, amount.ToFloat64())
}

func Test_should_be_format_money_to_specific_precision(t *testing.T) {
	testObjects := []*MoneyTestObject{
		{Expected: valueobjects.Money(10.2), Precision: 1, Value: 10.245678},
		{Expected: valueobjects.Money(10.25), Precision: 2, Value: 10.245678},
		{Expected: valueobjects.Money(10.246), Precision: 3, Value: 10.245678},
		{Expected: valueobjects.Money(10.2457), Precision: 4, Value: 10.245678},
	}

	for _, testObject := range testObjects {
		amount := valueobjects.NewMoney(testObject.Value)
		amount.Format(testObject.Precision)

		require.Equal(t, testObject.Expected, amount)
	}
}

func Test_should_be_return_an_error_when_amount_to_be_equal_zero(t *testing.T) {
	amount := valueobjects.NewMoney(0)

	err := amount.Validate()

	require.True(t, errors.As(err, &valueobjects.ErrorInvalidAmmount))
}
