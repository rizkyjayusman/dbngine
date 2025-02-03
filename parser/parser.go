package parser

import (
	"errors"
)

type TokenValidatorParam struct {
	pos int
}

type Parser struct {
	Tokens []Token
}

func NewParser(Tokens []Token) *Parser {
	return &Parser{
		Tokens: Tokens,
	}
}

func (p *Parser) Parse() (ASTNode, error) {
	if p.Tokens[0].Type != KEYWORD {
		return nil, errors.New("expected KEYWORD")
	}

	var node ASTNode
	var err error

	if p.Tokens[0].Value == SELECT {
		node, err = p.parseSelect(p.Tokens)
	}

	if err != nil {
		return nil, err
	}

	return node, nil
}

func (p *Parser) parseSelect(tokens []Token) (*SelectStatement, error) {
	param := TokenValidatorParam{pos: 0}

	if tokens[0].Type != KEYWORD && tokens[0].Value != SELECT {
		return &SelectStatement{}, nil
	}

	node := &SelectStatement{}

	param.pos++
	nextShouldDelimiter := false
	for param.pos < len(p.Tokens) {

		if p.Tokens[param.pos].Type == OPERATOR && p.Tokens[param.pos].Value == WILDCARD {
			node.Columns = append(node.Columns, p.Tokens[param.pos].Value)
			param.pos++
			break
		}

		if !nextShouldDelimiter && p.Tokens[param.pos].Type != IDENTIFIER {
			return node, errors.New("expected IDENTIFIER")
		}

		if p.Tokens[param.pos].Type != IDENTIFIER && p.Tokens[param.pos].Type != DELIMITER {
			break
		}

		if p.Tokens[param.pos].Type == IDENTIFIER {
			if nextShouldDelimiter {
				return node, errors.New("expected IDENTIFIER")
			}

			node.Columns = append(node.Columns, p.Tokens[param.pos].Value)
			nextShouldDelimiter = true
		} else {
			if !nextShouldDelimiter {
				return node, errors.New("expected DELIMITER")
			}

			nextShouldDelimiter = false
		}

		param.pos++
	}

	if p.Tokens[param.pos].Type == KEYWORD && p.Tokens[param.pos].Value == FROM {
		param.pos++
		if p.Tokens[param.pos].Type != IDENTIFIER {
			return node, errors.New("expected Table Name")
		}

		node.Table = p.Tokens[param.pos].Value
		param.pos++
	}

	whereClause, err := p.parseWhere(&param)
	node.WhereClause = whereClause

	if err != nil {
		return node, err
	}

	if param.pos < len(p.Tokens) {
		if p.Tokens[param.pos].Type != SYMBOL && p.Tokens[param.pos].Value != ";" {
			return node, errors.New("expected SYMBOL")
		}

		param.pos++
	}

	if param.pos == len(p.Tokens) {
		return node, nil
	}

	return node, errors.New("expected EOF")
}

func (p *Parser) ParseWhere() (*WhereClause, error) {
	pos := 0
	var root *WhereClause

	if pos >= len(p.Tokens) {
		return root, nil
	}

	if p.Tokens[pos].Type != KEYWORD && p.Tokens[pos].Value != WHERE {
		return root, nil
	}

	pos++
	
	root = p.parseWhereByToken(p.Tokens[pos:])

	pos++

	return root, nil
}

func (p *Parser) parseWhereByToken(tokens []Token) *WhereClause {
	if len(tokens) == 3 {
		return &WhereClause{
			Left:  &WhereClause{Name: tokens[0].Value},
			Type:  tokens[1].Value,
			Right: &WhereClause{Value: tokens[2].Value},
		}
	}

	return &WhereClause{
		Type:  tokens[3].Value,
		Left:  p.parseWhereByToken(tokens[:3]),
		Right: p.parseWhereByToken(tokens[4:]),
	}
}

func (p *Parser) parseWhere(param *TokenValidatorParam) (WhereClause, error) {
	root := WhereClause{}

	if param.pos < len(p.Tokens) && p.Tokens[param.pos].Type == KEYWORD && p.Tokens[param.pos].Value == WHERE {
		param.pos++

		var currentClause *WhereClause
		for param.pos < len(p.Tokens) {
			if p.Tokens[param.pos].Type != IDENTIFIER {
				return root, errors.New("expected IDENTIFIER")
			}

			newClause := &WhereClause{Name: p.Tokens[param.pos].Value}
			if currentClause == nil {
				currentClause.Right = newClause
			}
			param.pos++

			if p.Tokens[param.pos].Type != OPERATOR && p.Tokens[param.pos].Value != EQUALS {
				return root, errors.New("expected EQUALS")
			}

			root.Type = p.Tokens[param.pos].Value
			param.pos++

			if p.Tokens[param.pos].Type != LITERAL {
				return root, errors.New("expected LITERAL")
			}

			root.Right = &WhereClause{Name: p.Tokens[param.pos].Value}
			param.pos++

			if param.pos < len(p.Tokens) {
				if p.Tokens[param.pos].Type == OPERATOR {
					if p.Tokens[param.pos].Value != AND && p.Tokens[param.pos].Value != OR {
						break
					}

					param.pos++
				} else {
					break
				}
			}
		}
	}

	return root, nil
}

func (p *Parser) ValidateTokens() bool {
	param := TokenValidatorParam{pos: 0}

	if p.Tokens[param.pos].Type == KEYWORD {
		if p.Tokens[param.pos].Value == SELECT {
			return p.validateSelectTokens(param)
		} else if p.Tokens[param.pos].Value == INSERT {
			return p.validateInsertTokens(param)
		} else if p.Tokens[param.pos].Value == UPDATE {
			return p.validateUpdateTokens(param)
		} else if p.Tokens[param.pos].Value == DELETE {
			return p.validateDeleteTokens(param)
		}
	}
	return false
}

func (p *Parser) validateSelectTokens(param TokenValidatorParam) bool {
	param.pos++
	nextShouldDelimiter := false
	for param.pos < len(p.Tokens) {

		if p.Tokens[param.pos].Type == OPERATOR && p.Tokens[param.pos].Value == WILDCARD {
			param.pos++
			break
		}

		if !nextShouldDelimiter && p.Tokens[param.pos].Type != IDENTIFIER {
			return false
		}

		if p.Tokens[param.pos].Type != IDENTIFIER && p.Tokens[param.pos].Type != DELIMITER {
			break
		}

		if p.Tokens[param.pos].Type == IDENTIFIER {
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

		param.pos++
	}

	if p.Tokens[param.pos].Type == KEYWORD && p.Tokens[param.pos].Value == FROM {
		param.pos++
		if p.Tokens[param.pos].Type != IDENTIFIER {
			return false
		}

		param.pos++
	}

	res := p.validateWhereTokens(&param)
	if !res {
		return res
	}

	if param.pos < len(p.Tokens) {
		if p.Tokens[param.pos].Type != SYMBOL && p.Tokens[param.pos].Value != ";" {
			return false
		}

		param.pos++
	}

	if param.pos == len(p.Tokens) {
		return true
	}

	return false
}

func (p *Parser) validateWhereTokens(param *TokenValidatorParam) bool {
	if param.pos < len(p.Tokens) && p.Tokens[param.pos].Type == KEYWORD && p.Tokens[param.pos].Value == WHERE {
		param.pos++

		for param.pos < len(p.Tokens) {
			if p.Tokens[param.pos].Type != IDENTIFIER {
				return false
			}

			param.pos++

			if p.Tokens[param.pos].Type != OPERATOR && p.Tokens[param.pos].Value != EQUALS {
				return false
			}

			param.pos++

			if p.Tokens[param.pos].Type != LITERAL {
				return false
			}

			param.pos++

			if param.pos < len(p.Tokens) {
				if p.Tokens[param.pos].Type == OPERATOR {
					if p.Tokens[param.pos].Value != AND && p.Tokens[param.pos].Value != OR {
						break
					}

					param.pos++
				} else {
					break
				}
			}
		}
	}

	return true
}

func (p *Parser) validateInsertTokens(param TokenValidatorParam) bool {
	param.pos++

	if p.Tokens[param.pos].Type != KEYWORD && p.Tokens[param.pos].Value != INTO {
		return false
	}

	param.pos++

	if p.Tokens[param.pos].Type != IDENTIFIER {
		return false
	}

	param.pos++

	if p.Tokens[param.pos].Type == SYMBOL && p.Tokens[param.pos].Value == "(" {
		param.pos++
		nextShouldDelimiter := false
		for param.pos < len(p.Tokens) {
			if p.Tokens[param.pos].Type == IDENTIFIER {
				if nextShouldDelimiter {
					return false
				}

				nextShouldDelimiter = true
				param.pos++
			} else if p.Tokens[param.pos].Type == DELIMITER {
				if !nextShouldDelimiter {
					return false
				}

				nextShouldDelimiter = false
				param.pos++
			} else if p.Tokens[param.pos].Value == ")" {
				param.pos++
				break
			} else {
				return false
			}
		}

		if p.Tokens[param.pos].Type != KEYWORD && p.Tokens[param.pos].Value != VALUES {
			return false
		}

		param.pos++

		if p.Tokens[param.pos].Type == SYMBOL && p.Tokens[param.pos].Value == "(" {
			param.pos++
			nextShouldDelimiter := false
			for param.pos < len(p.Tokens) {
				if p.Tokens[param.pos].Type == LITERAL {
					if nextShouldDelimiter {
						return false
					}

					nextShouldDelimiter = true
					param.pos++
				} else if p.Tokens[param.pos].Type == DELIMITER {
					if !nextShouldDelimiter {
						return false
					}

					nextShouldDelimiter = false
					param.pos++
				} else if p.Tokens[param.pos].Value == ")" {
					param.pos++
					break
				} else {
					return false
				}
			}

			if param.pos == len(p.Tokens) {
				return true
			}
		}
	}

	return false
}

func (p *Parser) validateUpdateTokens(param TokenValidatorParam) bool {
	param.pos++

	if p.Tokens[param.pos].Type != IDENTIFIER {
		return false
	}

	param.pos++

	if p.Tokens[param.pos].Type == KEYWORD && p.Tokens[param.pos].Value == SET {
		param.pos++
		for param.pos < len(p.Tokens) {
			if p.Tokens[param.pos].Type != IDENTIFIER {
				return false
			}

			param.pos++

			if p.Tokens[param.pos].Type == OPERATOR && p.Tokens[param.pos].Value != EQUALS {
				return false
			}

			param.pos++

			if p.Tokens[param.pos].Type != LITERAL {
				return false
			}

			param.pos++

			if p.Tokens[param.pos].Type == DELIMITER && p.Tokens[param.pos].Value == "," {
				param.pos++
			} else {
				break
			}
		}
	}

	if param.pos < len(p.Tokens) && p.Tokens[param.pos].Type == KEYWORD && p.Tokens[param.pos].Value == WHERE {
		param.pos++

		for param.pos < len(p.Tokens) {
			if p.Tokens[param.pos].Type != IDENTIFIER {
				return false
			}

			param.pos++

			if p.Tokens[param.pos].Type != OPERATOR && p.Tokens[param.pos].Value != EQUALS {
				return false
			}

			param.pos++

			if p.Tokens[param.pos].Type != LITERAL {
				return false
			}

			param.pos++

			if param.pos < len(p.Tokens) {
				if p.Tokens[param.pos].Type == OPERATOR {
					if p.Tokens[param.pos].Value != AND && p.Tokens[param.pos].Value != OR {
						break
					}

					param.pos++
				} else {
					break
				}
			}
		}
	}

	if param.pos == len(p.Tokens) {
		return true
	}

	return false
}

func (p *Parser) validateDeleteTokens(param TokenValidatorParam) bool {
	pos := 1

	if p.Tokens[pos].Type != KEYWORD && p.Tokens[pos].Value != FROM {
		return false
	}

	pos++

	if p.Tokens[pos].Type != IDENTIFIER {
		return false
	}

	pos++

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

	if pos == len(p.Tokens) {
		return true
	}

	return false
}
