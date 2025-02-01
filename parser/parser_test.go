package parser

import "testing"

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

func TestParser_ValidateTokens_SimpleSelectQuery(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: SELECT},
		{Type: IDENTIFIER, Value: "id"},
		{Type: KEYWORD, Value: FROM},
		{Type: IDENTIFIER, Value: "users"},
	}

	parser := NewParser(tokens)
	t.Run("Check validation of format query", func(t *testing.T) {
		if !parser.ValidateTokens() {
			t.Errorf("invalid formats.")
		}
	})
}

func TestParser_ValidateTokens_SelectQueryWithWildcard(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: SELECT},
		{Type: OPERATOR, Value: WILDCARD},
		{Type: KEYWORD, Value: FROM},
		{Type: IDENTIFIER, Value: "users"},
	}

	parser := NewParser(tokens)
	t.Run("Check validation of format query", func(t *testing.T) {
		if !parser.ValidateTokens() {
			t.Errorf("invalid formats.")
		}
	})
}

func TestParser_ValidateTokens_SelectQueryWithSemicolon(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: SELECT},
		{Type: IDENTIFIER, Value: "id"},
		{Type: KEYWORD, Value: FROM},
		{Type: IDENTIFIER, Value: "users"},
		{Type: SYMBOL, Value: ";"},
	}

	parser := NewParser(tokens)
	t.Run("Check validation of format query", func(t *testing.T) {
		if !parser.ValidateTokens() {
			t.Errorf("invalid formats.")
		}
	})
}

func TestParser_ValidateTokens_SelectQueryMultipleColumns(t *testing.T) {
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
	t.Run("Check validation of format query", func(t *testing.T) {
		if !parser.ValidateTokens() {
			t.Errorf("invalid formats.")
		}
	})
}

func TestParser_ValidateTokens_SelectQueryButCommaBeforeFromKeyword(t *testing.T) {
	tokens := []Token{
		{Type: KEYWORD, Value: SELECT},
		{Type: IDENTIFIER, Value: "id"},
		{Type: DELIMITER, Value: ","},
		{Type: KEYWORD, Value: FROM},
		{Type: IDENTIFIER, Value: "users"},
	}

	parser := NewParser(tokens)
	t.Run("Check validation of format query", func(t *testing.T) {
		if parser.ValidateTokens() {
			t.Errorf("invalid formats.")
		}
	})
}

func TestParser_ValidateTokens_SelectQueryWithWhereClause(t *testing.T) {
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
	t.Run("Check validation of format query", func(t *testing.T) {
		if !parser.ValidateTokens() {
			t.Errorf("invalid formats.")
		}
	})
}

func TestParser_ValidateTokens_SelectQueryWithMultipleAndWhereClauses(t *testing.T) {
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
	t.Run("Check validation of format query", func(t *testing.T) {
		if !parser.ValidateTokens() {
			t.Errorf("invalid formats.")
		}
	})
}

func TestParser_ValidateTokens_SelectQueryWithMultipleAndWhereClausesAndOrOperator(t *testing.T) {
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
	t.Run("Check validation of format query", func(t *testing.T) {
		if !parser.ValidateTokens() {
			t.Errorf("invalid formats.")
		}
	})
}

func TestParser_ValidateTokens_SimpleInsertQuery(t *testing.T) {
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
	t.Run("Check validation of format query", func(t *testing.T) {
		if !parser.ValidateTokens() {
			t.Errorf("invalid formats.")
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
	t.Run("Check validation of format query", func(t *testing.T) {
		if !parser.ValidateTokens() {
			t.Errorf("invalid formats.")
		}
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
	}

	parser := NewParser(tokens)
	t.Run("Check validation of format query", func(t *testing.T) {
		if !parser.ValidateTokens() {
			t.Errorf("invalid formats.")
		}
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
