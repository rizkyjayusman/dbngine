package parser

import (
	"dbngin3/engine"
	"testing"
)

func TestSelectStatement_Analyze(t *testing.T) {
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
		Columns: []string{"id", "name"},
	}

	err := selectStmt.Analyze(schema)
	t.Run("Analyze AST nodes using Semantic", func(t *testing.T) {
		if err != nil {
			t.Error(err)
		}
	})
}
