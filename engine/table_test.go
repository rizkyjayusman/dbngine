package engine

import "testing"

func TestNewTable(t *testing.T) {
	table := NewTable("t_users", []Column{
		{Name: "id", Type: Int},
		{Name: "name", Type: Varchar},
	})

	t.Run("check table instance", func(t *testing.T) {
		if table == nil {
			t.Fatal("table is nil")
		}

		if table.Name != "t_users" {
			t.Fatal("table name is wrong")
		}

		if len(table.Columns) != 2 {
			t.Error("columns is wrong")
		}

		if table.Columns[0].Name != "id" || table.Columns[1].Name != "name" {
			t.Error("column name is wrong")
		}

		if table.Columns[0].Type != Int || table.Columns[1].Type != Varchar {
			t.Error("column type is wrong")
		}
	})

}
