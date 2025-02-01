package parser

type TokenType int

const (
	KEYWORD TokenType = iota
	IDENTIFIER
	OPERATOR
	LITERAL
	DELIMITER
	SYMBOL
)

type KeywordType string

const (
	SELECT = "SELECT"
	FROM   = "FROM"
	WHERE  = "WHERE"
	INSERT = "INSERT"
	INTO   = "INTO"
	VALUES = "VALUES"
	UPDATE = "UPDATE"
	SET    = "SET"
	DELETE = "DELETE"
)

type OperatorType string

const (
	WILDCARD = "*"
	EQUALS   = "="
	AND      = "AND"
	OR       = "OR"
)

func GetKeywordOrIdentifier(value string) TokenType {
	switch value {
	case SELECT, FROM, WHERE, INSERT, INTO, VALUES, UPDATE, SET, DELETE:
		return KEYWORD
	}

	return IDENTIFIER
}

func IsConditionalOperator(str string) bool {
	switch str {
	case AND, OR:
		return true
	}
	return false
}

type Token struct {
	Type  TokenType
	Value string
}
