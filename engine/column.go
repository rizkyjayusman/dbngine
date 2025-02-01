package engine

type Column struct {
	Name string
	Type DataType
}

func NewColumn(name string, dataType DataType) *Column {
	return &Column{
		Name: name,
		Type: dataType,
	}
}
