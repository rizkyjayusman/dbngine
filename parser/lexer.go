package parser

import (
	"dbngin3/util"
	"fmt"
)

type Lexer struct {
	Input string
}

func InitLexer() *Lexer {
	return &Lexer{}
}

func (l *Lexer) SetInput(input string) error {
	l.Input = input
	return nil
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		Input: input,
	}
}

func (l *Lexer) Tokenize() ([]Token, error) {
	var tokens []Token

	input := l.Input
	pos := 0
	for pos < len(input) {
		char := input[pos]

		if util.IsWhitespace(char) {
			pos++
			continue
		}

		if util.IsLetter(char) {
			start := pos
			for pos < len(input) && (util.IsLetter(input[pos])) {
				pos++
			}
			value := input[start:pos]
			if IsConditionalOperator(value) {
				tokens = append(tokens, Token{Type: OPERATOR, Value: value})
				continue
			}
			tokenType := GetKeywordOrIdentifier(value)
			tokens = append(tokens, Token{Type: tokenType, Value: value})
			continue
		}

		if util.IsDigit(char) {
			start := pos
			for pos < len(input) && util.IsDigit(input[pos]) {
				pos++
			}
			value := input[start:pos]
			tokens = append(tokens, Token{Type: LITERAL, Value: value})
			continue
		}

		if char == '\'' || char == '"' {
			start := pos + 1
			pos++

			for pos < len(input) && !(input[pos] == '\'' || input[pos] == '"') {
				pos++
			}

			if pos < len(input) && (input[pos] == '\'' || input[pos] == '"') {
				value := input[start:pos]
				tokens = append(tokens, Token{Type: LITERAL, Value: value})
				pos++
			} else {
				return nil, fmt.Errorf("unclosed string literal")
			}
			continue
		}

		if util.IsOperator(char) {
			operator := string(char)
			if util.IsOperator(input[pos+1]) {
				operator = operator + string(input[pos+1])
				pos++
			}

			tokens = append(tokens, Token{Type: OPERATOR, Value: operator})
			pos++
			continue
		}

		if util.IsDelimiter(char) {
			tokens = append(tokens, Token{Type: DELIMITER, Value: string(char)})
			pos++
			continue
		}

		if util.IsSymbol(char) {
			tokens = append(tokens, Token{Type: SYMBOL, Value: string(char)})
			pos++
			continue
		}

		pos++
	}

	return tokens, nil
}
