package valueobjects

import "errors"

var ErrorInvalidOperation = errors.New(`"operation must be Debit(0) or Credit(1)"`)

const (
	Debit  Operation = 0
	Credit Operation = 1
)

type Operation byte

func NewOperation(operation Operation) Operation {
	return Operation(operation)
}

func (o *Operation) Validate() error {
	if *o > Credit {
		return ErrorInvalidOperation
	}

	return nil
}

func (o *Operation) IsDebit() bool {
	return *o == Debit
}

func (o *Operation) IsCredit() bool {
	return *o == Credit
}
