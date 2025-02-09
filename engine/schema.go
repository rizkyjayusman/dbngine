package engine

import (
	"dbngin3/storage"
	"encoding/json"
	"errors"
)

type Schema struct {
	Tables []*Table `json:"tables"`
}

type SchemaManager struct {
	tables map[string]*Table
}

func NewSchemaManager() *SchemaManager {
	storageObj, err := storage.Open("storage/schema.json")
	if err != nil {
		panic(err)
	}

	file := storageObj.Read()
	var schema Schema
	err = json.Unmarshal(file, &schema)
	if err != nil {
		panic(err)
	}
	err = storageObj.Close()
	if err != nil {
		return nil
	}

	tables := make(map[string]*Table)
	for _, table := range schema.Tables {
		tables[table.Name] = table
	}

	return &SchemaManager{
		tables: tables,
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
