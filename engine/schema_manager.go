package engine

import "errors"

type SchemaManager struct {
	tables map[string]*Table
}

func NewSchemaManager() *SchemaManager {
	return &SchemaManager{
		tables: make(map[string]*Table),
	}
}

func (sm *SchemaManager) AddTable(name string, table *Table) {
	sm.tables[name] = table
}

func (sm *SchemaManager) GetTable(name string) (*Table, error) {
	res, ok := sm.tables[name]
	if !ok {
		return nil, errors.New("table not found")
	}

	return res, nil
}

func (sm *SchemaManager) IsTableExists(name string) bool {
	_, ok := sm.tables[name]
	return ok
}
