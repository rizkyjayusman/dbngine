package parser

// ASTNode : Abstract Syntax Tree
type ASTNode interface{}

type SelectStatement struct {
	Columns     []string
	Table       string
	WhereClause *WhereClause
}

type WhereClause struct {
	Type  string
	Left  *WhereClause
	Right *WhereClause
	Name  string
	Value string
}
