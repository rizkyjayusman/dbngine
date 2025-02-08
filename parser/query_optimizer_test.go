package parser

import (
	"dbngin3/engine"
	"reflect"
	"testing"
)

func TestSelectStatement_Optimize_WildcardSelect(t *testing.T) {
	schema := engine.NewSchemaManager()
	schema.AddTable("users", &engine.Table{
		Name: "users",
		Columns: []engine.Column{
			{Name: "id", Type: engine.Int},
			{Name: "name", Type: engine.Varchar},
		},
	})

	selectStmt := &SelectStatement{
		Table:   "users",
		Columns: []string{"*"},
	}

	err := selectStmt.Optimize(schema)
	t.Run("Optimize AST nodes", func(t *testing.T) {
		if err != nil {
			t.Error(err)
		}

		curTable, _ := schema.GetTable("users")
		expectedColumns := make([]string, 0)
		for _, col := range curTable.Columns {
			expectedColumns = append(expectedColumns, col.Name)
		}

		if !reflect.DeepEqual(selectStmt.Columns, expectedColumns) {
			t.Errorf("expected: %v, got: %v", expectedColumns, selectStmt.Columns)
		}
	})
}
