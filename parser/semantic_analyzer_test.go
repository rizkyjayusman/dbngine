package parser

import (
	"dbngin3/engine"
	"testing"
)

func TestSelectStatement_Analyze_SimpleSelectQuery(t *testing.T) {
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

	selectSemanticAnalyzer := SelectSemanticAnalyzer{
		Schema: schema,
	}

	err := selectSemanticAnalyzer.Analyze(selectStmt)
	t.Run("Analyze AST nodes using Semantic", func(t *testing.T) {
		if err != nil {
			t.Error(err)
		}
	})
}

func TestSelectStatement_Analyze_SelectQueryWithWhereClause(t *testing.T) {
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
		WhereClause: &WhereClause{
			Type: EQUALS,
			Left: &WhereClause{
				Name: "id",
			},
			Right: &WhereClause{
				Value: "12",
			},
		},
	}

	selectSemanticAnalyzer := SelectSemanticAnalyzer{
		Schema: schema,
	}

	err := selectSemanticAnalyzer.Analyze(selectStmt)
	t.Run("Analyze AST nodes using Semantic", func(t *testing.T) {
		if err != nil {
			t.Error(err)
		}
	})
}
