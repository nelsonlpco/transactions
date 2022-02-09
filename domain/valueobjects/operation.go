package valueobjects

const (
	Debit  Operation = iota
	Credit Operation = iota
)

type Operation byte

func NewOperation(operation Operation) Operation {
	return Operation(operation)
}

func (o *Operation) IsValid() bool {
	return *o <= Credit
}

func (o *Operation) ToByte() byte {
	return byte(*o)
}

func (o *Operation) IsDebit() bool {
	return *o == Debit
}

func (o *Operation) IsCredit() bool {
	return *o == Credit
}
