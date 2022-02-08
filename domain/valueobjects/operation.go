package valueobjects

const (
	Debit  byte = iota
	Credit byte = iota
)

type Operation byte

func NewOperation(operation byte) Operation {
	return Operation(operation)
}

func (o *Operation) IsValid() bool {
	return o.ToByte() <= Credit
}

func (o *Operation) ToByte() byte {
	return byte(*o)
}

func (o *Operation) IsDebit() bool {
	return o.ToByte() == Debit
}

func (o *Operation) IsCredit() bool {
	return o.ToByte() == Credit
}
