package parser

import "testing"

type TokenTest struct {
	name     string
	current  Token
	expected Token
}

func TestLexer_Init(t *testing.T) {
	query := "SELECT * FROM users"
	lexer := NewLexer(query)

	t.Run("Check lexer is not nil", func(t *testing.T) {
		if lexer == nil {
			t.Fatalf("expected lexer instance, got nil")
		}
	})

	t.Run("Check input is assigned correctly", func(t *testing.T) {
		expected := "SELECT * FROM users"
		if lexer.Input != expected {
			t.Errorf("expected input %q, got %q", expected, lexer.Input)
		}
	})
}

func TestLexer_Tokenize_SimpleSelectQuery(t *testing.T) {
	query := "SELECT id FROM users"
	lexer := NewLexer(query)
	tokens, _ := lexer.Tokenize()

	t.Run("Check tokens generated correctly", func(t *testing.T) {
		if len(tokens) != 4 {
			t.Errorf("expected 4 tokens, got %v", len(tokens))
		}
	})

	tests := []TokenTest{
		{"Check token at index 0 generated correctly", tokens[0], Token{Type: KEYWORD, Value: SELECT}},
		{"Check token at index 1 generated correctly", tokens[1], Token{Type: IDENTIFIER, Value: "id"}},
		{"Check token at index 2 generated correctly", tokens[2], Token{Type: KEYWORD, Value: FROM}},
		{"Check token at index 3 generated correctly", tokens[3], Token{Type: IDENTIFIER, Value: "users"}},
	}
	validateTokenDetail(t, tests)
}

func TestLexer_Tokenize_SelectQueryWithWildcard(t *testing.T) {
	query := "SELECT * FROM users"
	lexer := NewLexer(query)
	tokens, _ := lexer.Tokenize()

	t.Run("Check tokens generated correctly", func(t *testing.T) {
		if len(tokens) != 4 {
			t.Errorf("expected 4 tokens, got %v", len(tokens))
		}
	})

	tests := []TokenTest{
		{"Check token at index 0 generated correctly", tokens[0], Token{Type: KEYWORD, Value: SELECT}},
		{"Check token at index 1 generated correctly", tokens[1], Token{Type: OPERATOR, Value: WILDCARD}},
		{"Check token at index 2 generated correctly", tokens[2], Token{Type: KEYWORD, Value: FROM}},
		{"Check token at index 3 generated correctly", tokens[3], Token{Type: IDENTIFIER, Value: "users"}},
	}
	validateTokenDetail(t, tests)
}

func TestLexer_Tokenize_WithSemicolon(t *testing.T) {
	lexer := NewLexer("SELECT id FROM users;")
	tokens, _ := lexer.Tokenize()

	t.Run("Check tokens generated correctly", func(t *testing.T) {
		if len(tokens) != 5 {
			t.Errorf("expected 5 tokens, got %v", len(tokens))
		}
	})

	tests := []TokenTest{
		{"Check token at index 0 generated correctly", tokens[0], Token{Type: KEYWORD, Value: SELECT}},
		{"Check token at index 1 generated correctly", tokens[1], Token{Type: IDENTIFIER, Value: "id"}},
		{"Check token at index 2 generated correctly", tokens[2], Token{Type: KEYWORD, Value: FROM}},
		{"Check token at index 3 generated correctly", tokens[3], Token{Type: IDENTIFIER, Value: "users"}},
		{"Check token at index 4 generated correctly", tokens[4], Token{Type: SYMBOL, Value: ";"}},
	}
	validateTokenDetail(t, tests)
}

func TestLexer_Tokenize_SelectQueryMultipleColumns(t *testing.T) {
	query := "SELECT id, name, age FROM users"
	lexer := NewLexer(query)
	tokens, _ := lexer.Tokenize()

	t.Run("Check tokens generated correctly", func(t *testing.T) {
		if len(tokens) != 8 {
			t.Errorf("expected 4 tokens, got %v", len(tokens))
		}
	})

	tests := []TokenTest{
		{"Check token at index 0 generated correctly", tokens[0], Token{Type: KEYWORD, Value: SELECT}},
		{"Check token at index 1 generated correctly", tokens[1], Token{Type: IDENTIFIER, Value: "id"}},
		{"Check token at index 1 generated correctly", tokens[2], Token{Type: DELIMITER, Value: ","}},
		{"Check token at index 1 generated correctly", tokens[3], Token{Type: IDENTIFIER, Value: "name"}},
		{"Check token at index 1 generated correctly", tokens[4], Token{Type: DELIMITER, Value: ","}},
		{"Check token at index 1 generated correctly", tokens[5], Token{Type: IDENTIFIER, Value: "age"}},
		{"Check token at index 2 generated correctly", tokens[6], Token{Type: KEYWORD, Value: FROM}},
		{"Check token at index 3 generated correctly", tokens[7], Token{Type: IDENTIFIER, Value: "users"}},
	}
	validateTokenDetail(t, tests)
}

func TestLexer_Tokenize_SelectQueryWithWhereClause(t *testing.T) {
	lexer := NewLexer("SELECT id FROM users WHERE id = 1")
	tokens, _ := lexer.Tokenize()

	t.Run("Check tokens generated correctly", func(t *testing.T) {
		if len(tokens) != 8 {
			t.Errorf("expected 8 tokens, got %v", len(tokens))
		}
	})

	tests := []TokenTest{
		{"Check token at index 0 generated correctly", tokens[0], Token{Type: KEYWORD, Value: SELECT}},
		{"Check token at index 1 generated correctly", tokens[1], Token{Type: IDENTIFIER, Value: "id"}},
		{"Check token at index 2 generated correctly", tokens[2], Token{Type: KEYWORD, Value: FROM}},
		{"Check token at index 3 generated correctly", tokens[3], Token{Type: IDENTIFIER, Value: "users"}},
		{"Check token at index 4 generated correctly", tokens[4], Token{Type: KEYWORD, Value: WHERE}},
		{"Check token at index 5 generated correctly", tokens[5], Token{Type: IDENTIFIER, Value: "id"}},
		{"Check token at index 6 generated correctly", tokens[6], Token{Type: OPERATOR, Value: EQUALS}},
		{"Check token at index 7 generated correctly", tokens[7], Token{Type: LITERAL, Value: "1"}},
	}
	validateTokenDetail(t, tests)
}

func TestLexer_Tokenize_SelectQueryWithMultipleAndWhereClauses(t *testing.T) {
	lexer := NewLexer("SELECT id FROM users WHERE name = \"marty\" AND age = 18")
	tokens, _ := lexer.Tokenize()

	t.Run("Check tokens generated correctly", func(t *testing.T) {
		if len(tokens) != 12 {
			t.Errorf("expected 12 tokens, got %v", len(tokens))
		}
	})

	tests := []TokenTest{
		{"Check token at index 0 generated correctly", tokens[0], Token{Type: KEYWORD, Value: SELECT}},
		{"Check token at index 1 generated correctly", tokens[1], Token{Type: IDENTIFIER, Value: "id"}},
		{"Check token at index 2 generated correctly", tokens[2], Token{Type: KEYWORD, Value: FROM}},
		{"Check token at index 3 generated correctly", tokens[3], Token{Type: IDENTIFIER, Value: "users"}},
		{"Check token at index 4 generated correctly", tokens[4], Token{Type: KEYWORD, Value: WHERE}},
		{"Check token at index 5 generated correctly", tokens[5], Token{Type: IDENTIFIER, Value: "name"}},
		{"Check token at index 6 generated correctly", tokens[6], Token{Type: OPERATOR, Value: EQUALS}},
		{"Check token at index 7 generated correctly", tokens[7], Token{Type: LITERAL, Value: "marty"}},
		{"Check token at index 8 generated correctly", tokens[8], Token{Type: OPERATOR, Value: AND}},
		{"Check token at index 9 generated correctly", tokens[9], Token{Type: IDENTIFIER, Value: "age"}},
		{"Check token at index 10 generated correctly", tokens[10], Token{Type: OPERATOR, Value: EQUALS}},
		{"Check token at index 11 generated correctly", tokens[11], Token{Type: LITERAL, Value: "18"}},
	}
	validateTokenDetail(t, tests)
}

func TestLexer_Tokenize_SelectQueryWithMultipleAndWhereClausesAndOrOperator(t *testing.T) {
	lexer := NewLexer("SELECT id FROM users WHERE name = \"marty\" OR age = 18")
	tokens, _ := lexer.Tokenize()

	t.Run("Check tokens generated correctly", func(t *testing.T) {
		if len(tokens) != 12 {
			t.Errorf("expected 12 tokens, got %v", len(tokens))
		}
	})

	tests := []TokenTest{
		{"Check token at index 0 generated correctly", tokens[0], Token{Type: KEYWORD, Value: SELECT}},
		{"Check token at index 1 generated correctly", tokens[1], Token{Type: IDENTIFIER, Value: "id"}},
		{"Check token at index 2 generated correctly", tokens[2], Token{Type: KEYWORD, Value: FROM}},
		{"Check token at index 3 generated correctly", tokens[3], Token{Type: IDENTIFIER, Value: "users"}},
		{"Check token at index 4 generated correctly", tokens[4], Token{Type: KEYWORD, Value: WHERE}},
		{"Check token at index 5 generated correctly", tokens[5], Token{Type: IDENTIFIER, Value: "name"}},
		{"Check token at index 6 generated correctly", tokens[6], Token{Type: OPERATOR, Value: EQUALS}},
		{"Check token at index 7 generated correctly", tokens[7], Token{Type: LITERAL, Value: "marty"}},
		{"Check token at index 8 generated correctly", tokens[8], Token{Type: OPERATOR, Value: OR}},
		{"Check token at index 9 generated correctly", tokens[9], Token{Type: IDENTIFIER, Value: "age"}},
		{"Check token at index 10 generated correctly", tokens[10], Token{Type: OPERATOR, Value: EQUALS}},
		{"Check token at index 11 generated correctly", tokens[11], Token{Type: LITERAL, Value: "18"}},
	}
	validateTokenDetail(t, tests)
}

func TestLexer_Tokenize_SimpleInsertQuery(t *testing.T) {
	lexer := NewLexer("INSERT INTO users (id, name, age) VALUES (1, \"marty\", 18)")
	tokens, _ := lexer.Tokenize()

	t.Run("Check tokens generated correctly", func(t *testing.T) {
		if len(tokens) != 18 {
			t.Errorf("expected 18 tokens, got %v", len(tokens))
		}
	})

	tests := []TokenTest{
		{"Check token at index 0 generated correctly", tokens[0], Token{Type: KEYWORD, Value: INSERT}},
		{"Check token at index 1 generated correctly", tokens[1], Token{Type: KEYWORD, Value: INTO}},
		{"Check token at index 2 generated correctly", tokens[2], Token{Type: IDENTIFIER, Value: "users"}},
		{"Check token at index 3 generated correctly", tokens[3], Token{Type: SYMBOL, Value: "("}},
		{"Check token at index 4 generated correctly", tokens[4], Token{Type: IDENTIFIER, Value: "id"}},
		{"Check token at index 5 generated correctly", tokens[5], Token{Type: DELIMITER, Value: ","}},
		{"Check token at index 6 generated correctly", tokens[6], Token{Type: IDENTIFIER, Value: "name"}},
		{"Check token at index 7 generated correctly", tokens[7], Token{Type: DELIMITER, Value: ","}},
		{"Check token at index 8 generated correctly", tokens[8], Token{Type: IDENTIFIER, Value: "age"}},
		{"Check token at index 9 generated correctly", tokens[9], Token{Type: SYMBOL, Value: ")"}},
		{"Check token at index 10 generated correctly", tokens[10], Token{Type: KEYWORD, Value: VALUES}},
		{"Check token at index 11 generated correctly", tokens[11], Token{Type: SYMBOL, Value: "("}},
		{"Check token at index 12 generated correctly", tokens[12], Token{Type: LITERAL, Value: "1"}},
		{"Check token at index 13 generated correctly", tokens[13], Token{Type: DELIMITER, Value: ","}},
		{"Check token at index 14 generated correctly", tokens[14], Token{Type: LITERAL, Value: "marty"}},
		{"Check token at index 15 generated correctly", tokens[15], Token{Type: DELIMITER, Value: ","}},
		{"Check token at index 16 generated correctly", tokens[16], Token{Type: LITERAL, Value: "18"}},
		{"Check token at index 17 generated correctly", tokens[17], Token{Type: SYMBOL, Value: ")"}},
	}
	validateTokenDetail(t, tests)
}

func TestLexer_Tokenize_SimpleUpdateQuery(t *testing.T) {
	lexer := NewLexer("UPDATE users SET name = \"marty\" WHERE id = 1")
	tokens, _ := lexer.Tokenize()

	t.Run("Check tokens generated correctly", func(t *testing.T) {
		if len(tokens) != 10 {
			t.Errorf("expected 10 tokens, got %v", len(tokens))
		}
	})

	tests := []TokenTest{
		{"Check token at index 0 generated correctly", tokens[0], Token{Type: KEYWORD, Value: UPDATE}},
		{"Check token at index 2 generated correctly", tokens[1], Token{Type: IDENTIFIER, Value: "users"}},
		{"Check token at index 2 generated correctly", tokens[2], Token{Type: KEYWORD, Value: SET}},
		{"Check token at index 3 generated correctly", tokens[3], Token{Type: IDENTIFIER, Value: "name"}},
		{"Check token at index 4 generated correctly", tokens[4], Token{Type: OPERATOR, Value: EQUALS}},
		{"Check token at index 5 generated correctly", tokens[5], Token{Type: LITERAL, Value: "marty"}},
		{"Check token at index 6 generated correctly", tokens[6], Token{Type: KEYWORD, Value: WHERE}},
		{"Check token at index 7 generated correctly", tokens[7], Token{Type: IDENTIFIER, Value: "id"}},
		{"Check token at index 8 generated correctly", tokens[8], Token{Type: OPERATOR, Value: "="}},
		{"Check token at index 9 generated correctly", tokens[9], Token{Type: LITERAL, Value: "1"}},
	}
	validateTokenDetail(t, tests)
}

func TestLexer_Tokenize_UpdateQueryMultiColumnChanges(t *testing.T) {
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

func TestLexer_Tokenize_UpdateQueryMultiColumnWhereClauses(t *testing.T) {
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

func TestLexer_Tokenize_SimpleDeleteQuery(t *testing.T) {
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

func TestLexer_Tokenize_DeleteQueryMultipleWhereClauses(t *testing.T) {
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

func validateTokenDetail(t *testing.T, tests []TokenTest) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.current.Type != tt.expected.Type {
				t.Errorf("expected token type %v, got %v", tt.expected.Type, tt.current.Type)
			}

			if tt.current.Value != tt.expected.Value {
				t.Errorf("expected token value %v, got %v", tt.expected.Value, tt.current.Value)
			}
		})
	}
}
