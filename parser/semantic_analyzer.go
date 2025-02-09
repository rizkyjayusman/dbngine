package parser

import (
	"dbngin3/engine"
	"errors"
)

type SemanticAnalyzer interface {
	Analyze(node *ASTNode) error
}

type SelectSemanticAnalyzer struct {
	Schema *engine.SchemaManager
}

func (s *SelectSemanticAnalyzer) Analyze(selectStmt *SelectStatement) error {
	table, err := s.Schema.GetTable(selectStmt.Table)
	if err != nil {
		return errors.New("table not found in schema ")
	}

	for i := range selectStmt.Columns {
		if selectStmt.Columns[i] == WILDCARD {
			break
		}
		
		if !containsColumn(table.Columns, selectStmt.Columns[i]) {
			return errors.New("column not found in table ")
		}
	}

	whereColumns := selectStmt.WhereClause.GetColumnNames()
	for i := range whereColumns {
		if !containsColumn(table.Columns, whereColumns[i]) {
			return errors.New("column not found in table for where clause")
		}
	}

	return nil
}

func containsColumn(columns []engine.Column, col string) bool {
	for _, c := range columns {
		if c.Name == col {
			return true
		}
	}
	return false
}
