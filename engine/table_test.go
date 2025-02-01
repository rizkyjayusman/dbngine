package engine

import "testing"

func TestNewTable(t *testing.T) {
	table := NewTable("", []Column{
		{Name: "id", Type: Int},
		{Name: "name", Type: Varchar},
	})

	if table == nil {
		t.Error("table is nil")
	} else if len(table.Columns) != 2 {
		t.Error("columns is wrong")
	} else if table.Columns[0].Name != "id" || table.Columns[1].Name != "name" {
		t.Error("column name is wrong")
	} else if table.Columns[0].Type != Int || table.Columns[1].Type != Varchar {
		t.Error("column type is wrong")
	}

}
