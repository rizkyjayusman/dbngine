package engine

import "testing"

func TestNewColumn_IntType(t *testing.T) {
	column := NewColumn("id", Int)
	if column == nil {
		t.Error("column is nil")
	} else if column.Name != "id" {
		t.Error("column name is wrong")
	} else if column.Type != Int {
		t.Error("column type is wrong")
	}
}

func TestNewColumn_VarcharType(t *testing.T) {
	column := NewColumn("name", Varchar)
	if column == nil {
		t.Error("column is nil")
	} else if column.Name != "name" {
		t.Error("column name is wrong")
	} else if column.Type != Varchar {
		t.Error("column type is wrong")
	}
}
