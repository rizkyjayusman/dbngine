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

func (w *WhereClause) GetColumnNames() []string {
	return getColumns(w)
}

func getColumns(clause *WhereClause) []string {
	res := make([]string, 0)
	if clause == nil {
		return res
	}

	if clause.Left != nil && len(clause.Left.Name) > 0 {
		return append(res, clause.Left.Name)
	}

	left := getColumns(clause.Left)
	right := getColumns(clause.Right)

	res = append(res, left...)
	res = append(res, right...)

	return res
}
