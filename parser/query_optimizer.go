package parser

import "dbngin3/engine"

type QueryOptimizer interface {
	Optimize(node *ASTNode)
}

type SelectQueryOptimizer struct {
	Schema *engine.SchemaManager
}

func (s *SelectQueryOptimizer) Optimize(selectStmt *SelectStatement) error {
	table, err := s.Schema.GetTable(selectStmt.Table)
	if err != nil {
		return err
	}

	if selectStmt.Columns[0] == WILDCARD {
		selectStmt.Columns = []string{}
		for i := range table.Columns {
			selectStmt.Columns = append(selectStmt.Columns, table.Columns[i].Name)
		}
	}

	return nil
}
