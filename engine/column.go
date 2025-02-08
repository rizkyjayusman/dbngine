package engine

type Column struct {
	Name string   `json:"name"`
	Type DataType `json:"type"`
}

func NewColumn(name string, dataType DataType) *Column {
	return &Column{
		Name: name,
		Type: dataType,
	}
}
