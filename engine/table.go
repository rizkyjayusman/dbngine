package engine

type Table struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

func NewTable(name string, columns []Column) *Table {
	return &Table{
		Name:    name,
		Columns: columns,
	}
}
