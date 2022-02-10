package valueobjects

type Id int

func NewId(id int) Id {
	return Id(id)
}

func (i *Id) IsValid() bool {
	return int(*i) > 0
}
