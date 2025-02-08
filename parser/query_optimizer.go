package parser

import "dbngin3/engine"

type QueryOptimizer struct{}

func (s *SelectStatement) Optimize(schema *engine.SchemaManager) error {
	table, err := schema.GetTable(s.Table)
	if err != nil {
		return err
	}

	if s.Columns[0] == WILDCARD {
		s.Columns = []string{}
		for i := range table.Columns {
			s.Columns = append(s.Columns, table.Columns[i].Name)
		}
	}

	return nil
}
