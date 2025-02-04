package parser

// ASTNode : Abstract Syntax Tree
type ASTNode interface{}

type SelectStatement struct {
	Columns     []string
	Table       string
	WhereClause *WhereClause
}

type InsertStatement struct {
	Table   string
	Columns []string
	Values  []string
}

type UpdateStatement struct {
	Table       string
	Set         map[string]string
	WhereClause *WhereClause
}

type WhereClause struct {
	Type  string
	Left  *WhereClause
	Right *WhereClause
	Name  string
	Value string
}
