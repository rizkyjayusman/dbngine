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

func TestParser_ValidateQuery_SimpleSelectQuery(t *testing.T) {
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

func TestParser_Tokenize_WithSemicolon(t *testing.T) {
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

func TestParser_Tokenize_SelectQueryMultipleColumns(t *testing.T) {
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

func TestParser_Tokenize_SelectQueryWithWhereClause(t *testing.T) {
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

func TestParser_Tokenize_SelectQueryWithMultipleAndWhereClauses(t *testing.T) {
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

func TestParser_Tokenize_SelectQueryWithMultipleAndWhereClausesAndOrOperator(t *testing.T) {
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

func TestParser_Tokenize_SimpleInsertQuery(t *testing.T) {
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

func TestParser_Tokenize_SimpleUpdateQuery(t *testing.T) {
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

func TestParser_Tokenize_UpdateQueryMultiColumnChanges(t *testing.T) {
	lexer := NewLexer("UPDATE users SET name = \"marty\", age = 18, email = \"marty.mcfly@thefuture.com\" WHERE id = 1")
	tokens, _ := lexer.Tokenize()

	t.Run("Check tokens generated correctly", func(t *testing.T) {
		if len(tokens) != 18 {
			t.Errorf("expected 18 tokens, got %v", len(tokens))
		}
	})

	tests := []TokenTest{
		{"Check token at index 0 generated correctly", tokens[0], Token{Type: KEYWORD, Value: UPDATE}},
		{"Check token at index 2 generated correctly", tokens[1], Token{Type: IDENTIFIER, Value: "users"}},
		{"Check token at index 2 generated correctly", tokens[2], Token{Type: KEYWORD, Value: SET}},
		{"Check token at index 3 generated correctly", tokens[3], Token{Type: IDENTIFIER, Value: "name"}},
		{"Check token at index 4 generated correctly", tokens[4], Token{Type: OPERATOR, Value: EQUALS}},
		{"Check token at index 5 generated correctly", tokens[5], Token{Type: LITERAL, Value: "marty"}},
		{"Check token at index 5 generated correctly", tokens[6], Token{Type: DELIMITER, Value: ","}},
		{"Check token at index 5 generated correctly", tokens[7], Token{Type: IDENTIFIER, Value: "age"}},
		{"Check token at index 5 generated correctly", tokens[8], Token{Type: OPERATOR, Value: EQUALS}},
		{"Check token at index 5 generated correctly", tokens[9], Token{Type: LITERAL, Value: "18"}},
		{"Check token at index 5 generated correctly", tokens[10], Token{Type: DELIMITER, Value: ","}},
		{"Check token at index 5 generated correctly", tokens[11], Token{Type: IDENTIFIER, Value: "email"}},
		{"Check token at index 5 generated correctly", tokens[12], Token{Type: OPERATOR, Value: EQUALS}},
		{"Check token at index 5 generated correctly", tokens[13], Token{Type: LITERAL, Value: "marty.mcfly@thefuture.com"}},
		{"Check token at index 6 generated correctly", tokens[14], Token{Type: KEYWORD, Value: WHERE}},
		{"Check token at index 7 generated correctly", tokens[15], Token{Type: IDENTIFIER, Value: "id"}},
		{"Check token at index 8 generated correctly", tokens[16], Token{Type: OPERATOR, Value: "="}},
		{"Check token at index 9 generated correctly", tokens[17], Token{Type: LITERAL, Value: "1"}},
	}
	validateTokenDetail(t, tests)
}

func TestParser_Tokenize_UpdateQueryMultiColumnWhereClauses(t *testing.T) {
	lexer := NewLexer("UPDATE users SET name = \"marty\", age = 18 WHERE id = 1 AND email = \"marty.mcfly@thefuture.com\"")
	tokens, _ := lexer.Tokenize()

	t.Run("Check tokens generated correctly", func(t *testing.T) {
		if len(tokens) != 18 {
			t.Errorf("expected 18 tokens, got %v", len(tokens))
		}
	})

	tests := []TokenTest{
		{"Check token at index 0 generated correctly", tokens[0], Token{Type: KEYWORD, Value: UPDATE}},
		{"Check token at index 1 generated correctly", tokens[1], Token{Type: IDENTIFIER, Value: "users"}},
		{"Check token at index 2 generated correctly", tokens[2], Token{Type: KEYWORD, Value: SET}},
		{"Check token at index 3 generated correctly", tokens[3], Token{Type: IDENTIFIER, Value: "name"}},
		{"Check token at index 4 generated correctly", tokens[4], Token{Type: OPERATOR, Value: EQUALS}},
		{"Check token at index 5 generated correctly", tokens[5], Token{Type: LITERAL, Value: "marty"}},
		{"Check token at index 6 generated correctly", tokens[6], Token{Type: DELIMITER, Value: ","}},
		{"Check token at index 7 generated correctly", tokens[7], Token{Type: IDENTIFIER, Value: "age"}},
		{"Check token at index 8 generated correctly", tokens[8], Token{Type: OPERATOR, Value: EQUALS}},
		{"Check token at index 9 generated correctly", tokens[9], Token{Type: LITERAL, Value: "18"}},
		{"Check token at index 10 generated correctly", tokens[10], Token{Type: KEYWORD, Value: WHERE}},
		{"Check token at index 11 generated correctly", tokens[11], Token{Type: IDENTIFIER, Value: "id"}},
		{"Check token at index 12 generated correctly", tokens[12], Token{Type: OPERATOR, Value: "="}},
		{"Check token at index 13 generated correctly", tokens[13], Token{Type: LITERAL, Value: "1"}},
		{"Check token at index 14 generated correctly", tokens[14], Token{Type: OPERATOR, Value: AND}},
		{"Check token at index 15 generated correctly", tokens[15], Token{Type: IDENTIFIER, Value: "email"}},
		{"Check token at index 16 generated correctly", tokens[16], Token{Type: OPERATOR, Value: EQUALS}},
		{"Check token at index 17 generated correctly", tokens[17], Token{Type: LITERAL, Value: "marty.mcfly@thefuture.com"}},
	}
	validateTokenDetail(t, tests)
}

func TestParser_Tokenize_SimpleDeleteQuery(t *testing.T) {
	lexer := NewLexer("DELETE FROM users WHERE id = 1")
	tokens, _ := lexer.Tokenize()

	t.Run("Check tokens generated correctly", func(t *testing.T) {
		if len(tokens) != 7 {
			t.Errorf("expected 7 tokens, got %v", len(tokens))
		}
	})

	tests := []TokenTest{
		{"Check token at index 0 generated correctly", tokens[0], Token{Type: KEYWORD, Value: DELETE}},
		{"Check token at index 2 generated correctly", tokens[1], Token{Type: KEYWORD, Value: FROM}},
		{"Check token at index 2 generated correctly", tokens[2], Token{Type: IDENTIFIER, Value: "users"}},
		{"Check token at index 3 generated correctly", tokens[3], Token{Type: KEYWORD, Value: WHERE}},
		{"Check token at index 4 generated correctly", tokens[4], Token{Type: IDENTIFIER, Value: "id"}},
		{"Check token at index 5 generated correctly", tokens[5], Token{Type: OPERATOR, Value: EQUALS}},
		{"Check token at index 6 generated correctly", tokens[6], Token{Type: LITERAL, Value: "1"}},
	}
	validateTokenDetail(t, tests)
}

func TestParser_Tokenize_DeleteQueryMultipleWhereClauses(t *testing.T) {
	lexer := NewLexer("DELETE FROM users WHERE id = 1 AND email = \"marty.mcfly@thefuture.com\"")
	tokens, _ := lexer.Tokenize()

	t.Run("Check tokens generated correctly", func(t *testing.T) {
		if len(tokens) != 11 {
			t.Errorf("expected 11 tokens, got %v", len(tokens))
		}
	})

	tests := []TokenTest{
		{"Check token at index 0 generated correctly", tokens[0], Token{Type: KEYWORD, Value: DELETE}},
		{"Check token at index 2 generated correctly", tokens[1], Token{Type: KEYWORD, Value: FROM}},
		{"Check token at index 2 generated correctly", tokens[2], Token{Type: IDENTIFIER, Value: "users"}},
		{"Check token at index 3 generated correctly", tokens[3], Token{Type: KEYWORD, Value: WHERE}},
		{"Check token at index 4 generated correctly", tokens[4], Token{Type: IDENTIFIER, Value: "id"}},
		{"Check token at index 5 generated correctly", tokens[5], Token{Type: OPERATOR, Value: EQUALS}},
		{"Check token at index 6 generated correctly", tokens[6], Token{Type: LITERAL, Value: "1"}},
		{"Check token at index 7 generated correctly", tokens[7], Token{Type: OPERATOR, Value: AND}},
		{"Check token at index 8 generated correctly", tokens[8], Token{Type: IDENTIFIER, Value: "email"}},
		{"Check token at index 9 generated correctly", tokens[9], Token{Type: OPERATOR, Value: EQUALS}},
		{"Check token at index 10 generated correctly", tokens[10], Token{Type: LITERAL, Value: "marty.mcfly@thefuture.com"}},
	}
	validateTokenDetail(t, tests)
}
