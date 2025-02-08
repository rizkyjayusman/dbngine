package parser

import (
	"dbngin3/engine"
	"errors"
)

func (s *SelectStatement) Analyze(schema *engine.SchemaManager) error {
	table, err := schema.GetTable(s.Table)
	if err != nil {
		return errors.New("table not found in schema ")
	}

	for i := range s.Columns {
		if !containsColumn(table.Columns, s.Columns[i]) {
			return errors.New("column not found in table ")
		}
	}

	//whereClause := schema.FindWhere(s.WhereClause)
	//if !whereClause {
	//	return errors.New("where clause not found in schema ")
	//}

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
