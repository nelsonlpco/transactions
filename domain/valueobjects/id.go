package valueobjects

type Id int

func NewId(id int) Id {
	return Id(id)
}

func (i *Id) IsValid() bool {
	return i.ToInt() > 0
}

func (i *Id) ToInt() int {
	return int(*i)
}
