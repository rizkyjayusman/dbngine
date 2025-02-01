package parser

type Parser struct {
	Tokens []Token
}

func NewParser(Tokens []Token) *Parser {
	return &Parser{
		Tokens: Tokens,
	}
}

func (p *Parser) ValidateTokens() bool {
	if p.Tokens[0].Type == KEYWORD {
		if p.Tokens[0].Value == SELECT {
			return p.validateSelectTokens()
		} else if p.Tokens[0].Value == INSERT {
			return p.validateInsertTokens()
		} else if p.Tokens[0].Value == UPDATE {
			return p.validateUpdateTokens()
		}
	}
	return false
}

func (p *Parser) validateSelectTokens() bool {
	pos := 1
	nextShouldDelimiter := false
	for pos < len(p.Tokens) {
		if p.Tokens[pos].Type != IDENTIFIER && p.Tokens[pos].Type != DELIMITER {
			break
		}

		if p.Tokens[pos].Type == IDENTIFIER {
			if nextShouldDelimiter {
				return false
			}

			nextShouldDelimiter = true
		} else {
			if !nextShouldDelimiter {
				return false
			}

			nextShouldDelimiter = false
		}

		pos++
	}

	if p.Tokens[pos].Type == KEYWORD && p.Tokens[pos].Value == FROM {
		pos++
		if p.Tokens[pos].Type != IDENTIFIER {
			return false
		}

		pos++
	}

	if pos < len(p.Tokens) && p.Tokens[pos].Type == KEYWORD && p.Tokens[pos].Value == WHERE {
		pos++

		for pos < len(p.Tokens) {
			if p.Tokens[pos].Type != IDENTIFIER {
				return false
			}

			pos++

			if p.Tokens[pos].Type != OPERATOR && p.Tokens[pos].Value != EQUALS {
				return false
			}

			pos++

			if p.Tokens[pos].Type != LITERAL {
				return false
			}

			pos++

			if pos < len(p.Tokens) {
				if p.Tokens[pos].Type == OPERATOR {
					if p.Tokens[pos].Value != AND && p.Tokens[pos].Value != OR {
						break
					}

					pos++
				} else {
					break
				}
			}
		}

	}

	if pos < len(p.Tokens) {
		if p.Tokens[pos].Type != SYMBOL && p.Tokens[pos].Value != ";" {
			return false
		}

		pos++
	}

	if pos == len(p.Tokens) {
		return true
	}

	return false
}

func (p *Parser) validateInsertTokens() bool {
	pos := 1

	if p.Tokens[pos].Type != KEYWORD && p.Tokens[pos].Value != INTO {
		return false
	}

	pos++

	if p.Tokens[pos].Type != IDENTIFIER {
		return false
	}

	pos++

	if p.Tokens[pos].Type == SYMBOL && p.Tokens[pos].Value == "(" {
		pos++
		nextShouldDelimiter := false
		for pos < len(p.Tokens) {
			if p.Tokens[pos].Type == IDENTIFIER {
				if nextShouldDelimiter {
					return false
				}

				nextShouldDelimiter = true
				pos++
			} else if p.Tokens[pos].Type == DELIMITER {
				if !nextShouldDelimiter {
					return false
				}

				nextShouldDelimiter = false
				pos++
			} else if p.Tokens[pos].Value == ")" {
				pos++
				break
			} else {
				return false
			}
		}

		if p.Tokens[pos].Type != KEYWORD && p.Tokens[pos].Value != VALUES {
			return false
		}

		pos++

		if p.Tokens[pos].Type == SYMBOL && p.Tokens[pos].Value == "(" {
			pos++
			nextShouldDelimiter := false
			for pos < len(p.Tokens) {
				if p.Tokens[pos].Type == LITERAL {
					if nextShouldDelimiter {
						return false
					}

					nextShouldDelimiter = true
					pos++
				} else if p.Tokens[pos].Type == DELIMITER {
					if !nextShouldDelimiter {
						return false
					}

					nextShouldDelimiter = false
					pos++
				} else if p.Tokens[pos].Value == ")" {
					pos++
					break
				} else {
					return false
				}
			}

			if pos == len(p.Tokens) {
				return true
			}
		}
	}

	return false
}

func (p *Parser) validateUpdateTokens() bool {
	return true
}
