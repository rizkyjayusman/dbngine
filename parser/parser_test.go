package parser

import (
	"reflect"
	"testing"
)

func TestParser_Init(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: SELECT},
		{Type: IDENTIFIER, Value: "*"},
		{Type: KEYWORD, Value: FROM},
		{Type: IDENTIFIER, Value: "users"},
	}

	parser := NewParser(tokens)

	t.Run("Check parser is not nil", func(t *testing.T) {
		if parser == nil {
			t.Fatalf("expected parser instance, got nil")
		}
	})

	t.Run("Check input is assigned correctly", func(t *testing.T) {
		if len(parser.Tokens) != 4 {
			t.Errorf("expected tokens length %v, got %v", 4, len(parser.Tokens))
		}
	})
}

func TestParser_Parse_SimpleSelectQuery(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: SELECT},
		{Type: IDENTIFIER, Value: "id"},
		{Type: KEYWORD, Value: FROM},
		{Type: IDENTIFIER, Value: "users"},
	}

	parser := NewParser(tokens)
	node, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser parse failed: %v", err)
	}

	selectStmt, ok := node.(*SelectStatement)
	if !ok {
		t.Errorf("Expected ASTNode to be of type *SelectStatement, but got %v", reflect.TypeOf(node))
		return
	}

	t.Run("Check generated AST Nodes", func(t *testing.T) {
		expectedColumns := []string{"id"}
		if !reflect.DeepEqual(selectStmt.Columns, expectedColumns) {
			t.Errorf("expected column %v, got %v", expectedColumns, selectStmt.Columns)
		}

		expectedTable := "users"
		if selectStmt.Table != expectedTable {
			t.Errorf("expected table %v, got %v", expectedTable, selectStmt.Table)
		}
	})
}

func TestParser_Parse_SelectQueryWithWildcard(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: SELECT},
		{Type: OPERATOR, Value: WILDCARD},
		{Type: KEYWORD, Value: FROM},
		{Type: IDENTIFIER, Value: "users"},
	}

	parser := NewParser(tokens)
	node, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser parse failed: %v", err)
	}

	selectStmt, ok := node.(*SelectStatement)
	if !ok {
		t.Errorf("Expected ASTNode to be of type *SelectStatement, but got %v", reflect.TypeOf(node))
		return
	}

	t.Run("Check generated AST Nodes", func(t *testing.T) {
		expectedColumns := []string{WILDCARD}
		if !reflect.DeepEqual(selectStmt.Columns, expectedColumns) {
			t.Errorf("expected column %v, got %v", expectedColumns, selectStmt.Columns)
		}

		expectedTable := "users"
		if selectStmt.Table != expectedTable {
			t.Errorf("expected table %v, got %v", expectedTable, selectStmt.Table)
		}
	})
}

func TestParser_Parse_SelectQueryWithSemicolon(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: SELECT},
		{Type: IDENTIFIER, Value: "id"},
		{Type: KEYWORD, Value: FROM},
		{Type: IDENTIFIER, Value: "users"},
		{Type: SYMBOL, Value: ";"},
	}

	parser := NewParser(tokens)
	node, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser parse failed: %v", err)
	}

	selectStmt, ok := node.(*SelectStatement)
	if !ok {
		t.Errorf("Expected ASTNode to be of type *SelectStatement, but got %v", reflect.TypeOf(node))
		return
	}

	t.Run("Check generated AST Nodes", func(t *testing.T) {
		expectedColumns := []string{"id"}
		if !reflect.DeepEqual(selectStmt.Columns, expectedColumns) {
			t.Errorf("expected column %v, got %v", expectedColumns, selectStmt.Columns)
		}

		expectedTable := "users"
		if selectStmt.Table != expectedTable {
			t.Errorf("expected table %v, got %v", expectedTable, selectStmt.Table)
		}
	})
}

func TestParser_Parse_SelectQueryMultipleColumns(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: SELECT},
		{Type: IDENTIFIER, Value: "id"},
		{Type: DELIMITER, Value: ","},
		{Type: IDENTIFIER, Value: "name"},
		{Type: DELIMITER, Value: ","},
		{Type: IDENTIFIER, Value: "age"},
		{Type: KEYWORD, Value: FROM},
		{Type: IDENTIFIER, Value: "users"},
	}

	parser := NewParser(tokens)
	node, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser parse failed: %v", err)
	}

	selectStmt, ok := node.(*SelectStatement)
	if !ok {
		t.Errorf("Expected ASTNode to be of type *SelectStatement, but got %v", reflect.TypeOf(node))
		return
	}

	t.Run("Check generated AST Nodes", func(t *testing.T) {
		expectedColumns := []string{"id", "name", "age"}
		if !reflect.DeepEqual(selectStmt.Columns, expectedColumns) {
			t.Errorf("expected column %v, got %v", expectedColumns, selectStmt.Columns)
		}

		expectedTable := "users"
		if selectStmt.Table != expectedTable {
			t.Errorf("expected table %v, got %v", expectedTable, selectStmt.Table)
		}
	})
}

func TestParser_Parse_SelectQueryButCommaBeforeFromKeyword(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: SELECT},
		{Type: IDENTIFIER, Value: "id"},
		{Type: DELIMITER, Value: ","},
		{Type: KEYWORD, Value: FROM},
		{Type: IDENTIFIER, Value: "users"},
	}

	parser := NewParser(tokens)
	_, err := parser.Parse()
	t.Run("Check validation of format query", func(t *testing.T) {
		if err == nil {
			t.Error("expected error : invalid comma, got none")
		}
	})
}

func TestParser_Parse_SelectQueryWithWhereClause(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: SELECT},
		{Type: IDENTIFIER, Value: "id"},
		{Type: KEYWORD, Value: FROM},
		{Type: IDENTIFIER, Value: "users"},
		{Type: KEYWORD, Value: WHERE},
		{Type: IDENTIFIER, Value: "id"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "1"},
	}

	parser := NewParser(tokens)
	node, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser parse failed: %v", err)
	}

	selectStmt, ok := node.(*SelectStatement)
	if !ok {
		t.Errorf("Expected ASTNode to be of type *SelectStatement, but got %v", reflect.TypeOf(node))
		return
	}

	t.Run("Check generated AST Nodes", func(t *testing.T) {
		expectedColumns := []string{"id"}
		if !reflect.DeepEqual(selectStmt.Columns, expectedColumns) {
			t.Errorf("expected column %v, got %v", expectedColumns, selectStmt.Columns)
		}

		expectedTable := "users"
		if selectStmt.Table != expectedTable {
			t.Errorf("expected table %v, got %v", expectedTable, selectStmt.Table)
		}

		whereClauseTest := WhereClause{
			Type:  EQUALS,
			Left:  &WhereClause{Name: "id"},
			Right: &WhereClause{Value: "1"},
		}

		validateWhereNode(whereClauseTest, *selectStmt.WhereClause, t)
	})
}

func TestParser_Parse_SelectQueryWithMultipleAndWhereClauses(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: SELECT},
		{Type: IDENTIFIER, Value: "id"},
		{Type: KEYWORD, Value: FROM},
		{Type: IDENTIFIER, Value: "users"},
		{Type: KEYWORD, Value: WHERE},
		{Type: IDENTIFIER, Value: "name"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "marty"},
		{Type: OPERATOR, Value: AND},
		{Type: IDENTIFIER, Value: "age"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "18"},
	}

	parser := NewParser(tokens)
	node, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser parse failed: %v", err)
	}

	selectStmt, ok := node.(*SelectStatement)
	if !ok {
		t.Errorf("Expected ASTNode to be of type *SelectStatement, but got %v", reflect.TypeOf(node))
		return
	}

	t.Run("Check generated AST Nodes", func(t *testing.T) {
		expectedColumns := []string{"id"}
		if !reflect.DeepEqual(selectStmt.Columns, expectedColumns) {
			t.Errorf("expected column %v, got %v", expectedColumns, selectStmt.Columns)
		}

		expectedTable := "users"
		if selectStmt.Table != expectedTable {
			t.Errorf("expected table %v, got %v", expectedTable, selectStmt.Table)
		}

		whereClauseTests := WhereClause{
			Type: AND,
			Left: &WhereClause{
				Type:  EQUALS,
				Left:  &WhereClause{Name: "id"},
				Right: &WhereClause{Name: "1"},
			},
			Right: &WhereClause{
				Type:  EQUALS,
				Left:  &WhereClause{Name: "age"},
				Right: &WhereClause{Name: "18"},
			},
		}

		validateWhereNode(whereClauseTests, *selectStmt.WhereClause, t)
	})
}

func TestParser_Parse_SelectQueryWithWhereClauseOrOperator(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: SELECT},
		{Type: IDENTIFIER, Value: "id"},
		{Type: KEYWORD, Value: FROM},
		{Type: IDENTIFIER, Value: "users"},
		{Type: KEYWORD, Value: WHERE},
		{Type: IDENTIFIER, Value: "name"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "marty"},
		{Type: OPERATOR, Value: OR},
		{Type: IDENTIFIER, Value: "age"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "18"},
	}

	parser := NewParser(tokens)
	node, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser parse failed: %v", err)
	}

	selectStmt, ok := node.(*SelectStatement)
	if !ok {
		t.Errorf("Expected ASTNode to be of type *SelectStatement, but got %v", reflect.TypeOf(node))
		return
	}

	t.Run("Check generated AST Nodes", func(t *testing.T) {
		expectedColumns := []string{"id"}
		if !reflect.DeepEqual(selectStmt.Columns, expectedColumns) {
			t.Errorf("expected column %v, got %v", expectedColumns, selectStmt.Columns)
		}

		expectedTable := "users"
		if selectStmt.Table != expectedTable {
			t.Errorf("expected table %v, got %v", expectedTable, selectStmt.Table)
		}

		whereClauseTests := WhereClause{
			Type: OR,
			Left: &WhereClause{
				Type:  EQUALS,
				Left:  &WhereClause{Name: "name"},
				Right: &WhereClause{Name: "marty"},
			},
			Right: &WhereClause{
				Type:  EQUALS,
				Left:  &WhereClause{Name: "age"},
				Right: &WhereClause{Name: "18"},
			},
		}

		validateWhereNode(whereClauseTests, *selectStmt.WhereClause, t)
	})
}

func TestParser_Parse_SimpleInsertQuery(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: INSERT},
		{Type: KEYWORD, Value: INTO},
		{Type: IDENTIFIER, Value: "users"},
		{Type: SYMBOL, Value: "("},
		{Type: IDENTIFIER, Value: "id"},
		{Type: DELIMITER, Value: ","},
		{Type: IDENTIFIER, Value: "name"},
		{Type: DELIMITER, Value: ","},
		{Type: IDENTIFIER, Value: "age"},
		{Type: SYMBOL, Value: ")"},
		{Type: KEYWORD, Value: VALUES},
		{Type: SYMBOL, Value: "("},
		{Type: LITERAL, Value: "1"},
		{Type: DELIMITER, Value: ","},
		{Type: LITERAL, Value: "marty"},
		{Type: DELIMITER, Value: ","},
		{Type: LITERAL, Value: "18"},
		{Type: SYMBOL, Value: ")"},
	}

	parser := NewParser(tokens)
	node, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser parse failed: %v", err)
	}

	insertStmt, ok := node.(*InsertStatement)
	if !ok {
		t.Errorf("Expected ASTNode to be of type *InserStatement, but got %v", reflect.TypeOf(node))
		return
	}
	t.Run("Check generated AST Nodes", func(t *testing.T) {
		expectedColumns := []string{"id", "name", "age"}
		if !reflect.DeepEqual(insertStmt.Columns, expectedColumns) {
			t.Errorf("expected column %v, got %v", expectedColumns, insertStmt.Columns)
		}

		expectedTable := "users"
		if insertStmt.Table != expectedTable {
			t.Errorf("expected table %v, got %v", expectedTable, insertStmt.Table)
		}

		expectedValues := []string{"1", "marty", "18"}
		if !reflect.DeepEqual(insertStmt.Values, expectedValues) {
			t.Errorf("expected column %v, got %v", expectedValues, insertStmt.Values)
		}
	})
}

func TestParser_ValidateTokens_SimpleUpdateQuery(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: UPDATE},
		{Type: IDENTIFIER, Value: "users"},
		{Type: KEYWORD, Value: SET},
		{Type: IDENTIFIER, Value: "name"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "marty"},
		{Type: KEYWORD, Value: WHERE},
		{Type: IDENTIFIER, Value: "id"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "1"},
	}

	parser := NewParser(tokens)
	node, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser parse failed: %v", err)
	}

	updateStmt, ok := node.(*UpdateStatement)
	if !ok {
		t.Errorf("Expected ASTNode to be of type *UpdateStatement, but got %v", reflect.TypeOf(node))
		return
	}
	t.Run("Check generated AST Nodes", func(t *testing.T) {
		expectedTable := "users"
		if updateStmt.Table != expectedTable {
			t.Errorf("expected table %v, got %v", expectedTable, updateStmt.Table)
		}

		expectedSets := map[string]string{
			"name": "marty",
		}
		if !reflect.DeepEqual(updateStmt.Set, expectedSets) {
			t.Errorf("expected column %v, got %v", expectedSets, updateStmt.Set)
		}

		whereClauseTests := WhereClause{
			Type:  EQUALS,
			Left:  &WhereClause{Name: "id"},
			Right: &WhereClause{Value: "1"},
		}

		validateWhereNode(whereClauseTests, *updateStmt.WhereClause, t)
	})
}

func TestParser_ValidateTokens_UpdateQueryMultiColumnChanges(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: UPDATE},
		{Type: IDENTIFIER, Value: "users"},
		{Type: KEYWORD, Value: SET},
		{Type: IDENTIFIER, Value: "name"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "marty"},
		{Type: DELIMITER, Value: ","},
		{Type: IDENTIFIER, Value: "email"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "marty.mcfly@thefuture.com"},
		{Type: KEYWORD, Value: WHERE},
		{Type: IDENTIFIER, Value: "id"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "1"},
		{Type: OPERATOR, Value: AND},
		{Type: IDENTIFIER, Value: "age"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "18"},
	}

	parser := NewParser(tokens)
	node, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser parse failed: %v", err)
	}

	updateStmt, ok := node.(*UpdateStatement)
	if !ok {
		t.Errorf("Expected ASTNode to be of type *UpdateStatement, but got %v", reflect.TypeOf(node))
		return
	}
	t.Run("Check generated AST Nodes", func(t *testing.T) {
		expectedTable := "users"
		if updateStmt.Table != expectedTable {
			t.Errorf("expected table %v, got %v", expectedTable, updateStmt.Table)
		}

		expectedSets := map[string]string{
			"name":  "marty",
			"email": "marty.mcfly@thefuture.com",
		}
		if !reflect.DeepEqual(updateStmt.Set, expectedSets) {
			t.Errorf("expected column %v, got %v", expectedSets, updateStmt.Set)
		}

		whereClauseTests := WhereClause{
			Type: AND,
			Left: &WhereClause{
				Type:  EQUALS,
				Left:  &WhereClause{Name: "id"},
				Right: &WhereClause{Name: "1"},
			},
			Right: &WhereClause{
				Type:  EQUALS,
				Left:  &WhereClause{Name: "age"},
				Right: &WhereClause{Name: "18"},
			},
		}

		validateWhereNode(whereClauseTests, *updateStmt.WhereClause, t)
	})
}

func TestParser_ValidateTokens_UpdateQueryMultiColumnWhereClauses(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: UPDATE},
		{Type: IDENTIFIER, Value: "users"},
		{Type: KEYWORD, Value: SET},
		{Type: IDENTIFIER, Value: "name"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "marty"},
		{Type: DELIMITER, Value: ","},
		{Type: IDENTIFIER, Value: "age"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "18"},
		{Type: KEYWORD, Value: WHERE},
		{Type: IDENTIFIER, Value: "id"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "1"},
		{Type: OPERATOR, Value: AND},
		{Type: IDENTIFIER, Value: "email"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "marty.mcfly@thefuture.com"},
	}

	parser := NewParser(tokens)
	t.Run("Check validation of format query", func(t *testing.T) {
		if !parser.ValidateTokens() {
			t.Errorf("invalid formats.")
		}
	})
}

func TestParser_ValidateTokens_SimpleDeleteQuery(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: DELETE},
		{Type: KEYWORD, Value: FROM},
		{Type: IDENTIFIER, Value: "users"},
		{Type: KEYWORD, Value: WHERE},
		{Type: IDENTIFIER, Value: "id"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "1"},
	}

	parser := NewParser(tokens)
	t.Run("Check validation of format query", func(t *testing.T) {
		if !parser.ValidateTokens() {
			t.Errorf("invalid formats.")
		}
	})
}

func TestParser_ValidateTokens_DeleteQueryMultipleWhereClauses(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: DELETE},
		{Type: KEYWORD, Value: FROM},
		{Type: IDENTIFIER, Value: "users"},
		{Type: KEYWORD, Value: WHERE},
		{Type: IDENTIFIER, Value: "id"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "1"},
		{Type: OPERATOR, Value: AND},
		{Type: IDENTIFIER, Value: "email"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "marty.mcfly@thefuture.com"},
	}

	parser := NewParser(tokens)
	t.Run("Check validation of format query", func(t *testing.T) {
		if !parser.ValidateTokens() {
			t.Errorf("invalid formats.")
		}
	})
}

func TestParser_ParseWhere_Equals(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: WHERE},
		{Type: IDENTIFIER, Value: "id"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "1"},
	}

	parser := NewParser(tokens)
	node, _ := parser.ParseWhere(&TokenValidatorParam{pos: 0})

	t.Run("Check generated AST Nodes", func(t *testing.T) {
		whereClauseTest := WhereClause{
			Type:  EQUALS,
			Left:  &WhereClause{Name: "id"},
			Right: &WhereClause{Value: "1"},
		}

		validateWhereNode(whereClauseTest, *node, t)
	})
}

func TestParser_ParseWhere_MultipleEqualsWithAndOperator(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: WHERE},
		{Type: IDENTIFIER, Value: "id"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "1"},
		{Type: OPERATOR, Value: AND},
		{Type: IDENTIFIER, Value: "age"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "18"},
	}

	parser := NewParser(tokens)
	node, _ := parser.ParseWhere(&TokenValidatorParam{pos: 0})

	t.Run("Check generated AST Nodes", func(t *testing.T) {
		whereClauseTest := WhereClause{
			Type: AND,
			Left: &WhereClause{
				Type:  EQUALS,
				Left:  &WhereClause{Name: "id"},
				Right: &WhereClause{Value: "1"},
			},
			Right: &WhereClause{
				Type:  EQUALS,
				Left:  &WhereClause{Name: "age"},
				Right: &WhereClause{Value: "18"},
			},
		}

		validateWhereNode(whereClauseTest, *node, t)
	})
}

func TestParser_ParseWhere_MultipleEqualsWithOrOperator(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: WHERE},
		{Type: IDENTIFIER, Value: "id"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "1"},
		{Type: OPERATOR, Value: OR},
		{Type: IDENTIFIER, Value: "age"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "18"},
	}

	parser := NewParser(tokens)
	node, _ := parser.ParseWhere(&TokenValidatorParam{pos: 0})

	t.Run("Check generated AST Nodes", func(t *testing.T) {
		whereClauseTest := WhereClause{
			Type: OR,
			Left: &WhereClause{
				Type:  EQUALS,
				Left:  &WhereClause{Name: "id"},
				Right: &WhereClause{Value: "1"},
			},
			Right: &WhereClause{
				Type:  EQUALS,
				Left:  &WhereClause{Name: "age"},
				Right: &WhereClause{Value: "18"},
			},
		}

		validateWhereNode(whereClauseTest, *node, t)
	})
}

func TestParser_ParseWhere_WithMultipleAndOperator(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: WHERE},
		{Type: IDENTIFIER, Value: "id"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "1"},
		{Type: OPERATOR, Value: AND},
		{Type: IDENTIFIER, Value: "age"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "18"},
		{Type: OPERATOR, Value: AND},
		{Type: IDENTIFIER, Value: "email"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "marty.mcfly@thefuture.com"},
	}

	parser := NewParser(tokens)
	node, _ := parser.ParseWhere(&TokenValidatorParam{pos: 0})

	t.Run("Check generated AST Nodes", func(t *testing.T) {
		whereClauseTest := WhereClause{
			Type: AND,
			Left: &WhereClause{
				Type:  EQUALS,
				Left:  &WhereClause{Name: "id"},
				Right: &WhereClause{Value: "1"},
			},
			Right: &WhereClause{
				Type: AND,
				Left: &WhereClause{
					Type:  EQUALS,
					Left:  &WhereClause{Name: "age"},
					Right: &WhereClause{Value: "18"},
				},
				Right: &WhereClause{
					Type:  EQUALS,
					Left:  &WhereClause{Name: "email"},
					Right: &WhereClause{Value: "marty.mcfly@thefuture.com"},
				},
			},
		}

		validateWhereNode(whereClauseTest, *node, t)
	})
}

func TestParser_ParseWhere_WithMultipleOrOperator(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: WHERE},
		{Type: IDENTIFIER, Value: "id"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "1"},
		{Type: OPERATOR, Value: OR},
		{Type: IDENTIFIER, Value: "age"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "18"},
		{Type: OPERATOR, Value: OR},
		{Type: IDENTIFIER, Value: "email"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "marty.mcfly@thefuture.com"},
	}

	parser := NewParser(tokens)
	node, _ := parser.ParseWhere(&TokenValidatorParam{pos: 0})

	t.Run("Check generated AST Nodes", func(t *testing.T) {
		whereClauseTest := WhereClause{
			Type: OR,
			Left: &WhereClause{
				Type:  EQUALS,
				Left:  &WhereClause{Name: "id"},
				Right: &WhereClause{Value: "1"},
			},
			Right: &WhereClause{
				Type: OR,
				Left: &WhereClause{
					Type:  EQUALS,
					Left:  &WhereClause{Name: "age"},
					Right: &WhereClause{Value: "18"},
				},
				Right: &WhereClause{
					Type:  EQUALS,
					Left:  &WhereClause{Name: "email"},
					Right: &WhereClause{Value: "marty.mcfly@thefuture.com"},
				},
			},
		}

		validateWhereNode(whereClauseTest, *node, t)
	})
}

func TestParser_ParseWhere_WithAndOperatorBeforeOrOperator(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: WHERE},
		{Type: IDENTIFIER, Value: "id"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "1"},
		{Type: OPERATOR, Value: AND},
		{Type: IDENTIFIER, Value: "age"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "18"},
		{Type: OPERATOR, Value: OR},
		{Type: IDENTIFIER, Value: "email"},
		{Type: OPERATOR, Value: EQUALS},
		{Type: LITERAL, Value: "marty.mcfly@thefuture.com"},
	}

	parser := NewParser(tokens)
	node, _ := parser.ParseWhere(&TokenValidatorParam{pos: 0})

	t.Run("Check generated AST Nodes", func(t *testing.T) {
		whereClauseTest := WhereClause{
			Type: OR,
			Left: &WhereClause{
				Type: AND,
				Left: &WhereClause{
					Type:  EQUALS,
					Left:  &WhereClause{Name: "id"},
					Right: &WhereClause{Value: "1"},
				},
				Right: &WhereClause{
					Type:  EQUALS,
					Left:  &WhereClause{Name: "age"},
					Right: &WhereClause{Value: "18"},
				},
			},
			Right: &WhereClause{
				Type:  EQUALS,
				Left:  &WhereClause{Name: "email"},
				Right: &WhereClause{Value: "marty.mcfly@thefuture.com"},
			},
		}

		validateWhereNode(whereClauseTest, *node, t)
	})
}

func validateWhereNode(expected WhereClause, current WhereClause, t *testing.T) {
	if current.Left.Name != expected.Left.Name {
		t.Errorf("Check where clauses Column: expected %v, got %v", expected.Left.Name, current.Left.Name)
	}

	if current.Type != expected.Type {
		t.Errorf("Check where clauses Operator: expected %v, got %v", expected.Type, current.Type)
	}

	if current.Right.Value != expected.Right.Value {
		t.Errorf("Check where clauses Value: expected %v, got %v", expected.Right.Value, current.Right.Value)
	}

}
