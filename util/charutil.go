package util

import (
	"strings"
)

func IsWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func IsLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func IsDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func IsOperator(char byte) bool {
	operators := "+-*/=<>!"
	return strings.ContainsRune(operators, rune(char))
}

func IsDelimiter(char byte) bool {
	operators := ","
	return strings.ContainsRune(operators, rune(char))
}

func IsSymbol(char byte) bool {
	operators := "();"
	return strings.ContainsRune(operators, rune(char))
}
