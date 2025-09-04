package parser

import (
	"design-patterns/behavioural/interpreter/expressions"
	"fmt"
	"strconv"
)

type Parser struct {
	lexer *Lexer

	curToken  Token
	peekToken Token
}

func NewParser(lexer *Lexer) *Parser {
	p := &Parser{lexer: lexer}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseExpression() (expressions.IExpr, error) {
	return p.parseOrExpression()
}

func (p *Parser) parseOrExpression() (expressions.IExpr, error) {
	expr, err := p.parseAndExpression()
	if err != nil {
		return nil, err
	}

	for p.curToken.Type == OR {
		p.nextToken()
		right, err := p.parseAndExpression()
		if err != nil {
			return nil, err
		}
		expr = &expressions.Or{Left: expr, Right: right}
	}

	return expr, nil
}

func (p *Parser) parseAndExpression() (expressions.IExpr, error) {
	expr, err := p.parsePrimaryExpression()
	if err != nil {
		return nil, err
	}

	for p.curToken.Type == AND {
		p.nextToken()
		right, err := p.parsePrimaryExpression()
		if err != nil {
			return nil, err
		}
		expr = &expressions.And{Left: expr, Right: right}
	}

	return expr, nil
}

func (p *Parser) parsePrimaryExpression() (expressions.IExpr, error) {
	switch p.curToken.Type {
	case LPAREN:
		return p.parseGroupedExpression()
	case IDENTIFIER:
		return p.parseIdentifierExpression()
	default:
		return nil, fmt.Errorf("unexpected token: %s at position %d", p.curToken.Type, p.curToken.Pos)
	}
}

func (p *Parser) parseGroupedExpression() (expressions.IExpr, error) {
	p.nextToken()

	expr, err := p.parseOrExpression()
	if err != nil {
		return nil, err
	}

	if p.curToken.Type != RPAREN {
		return nil, fmt.Errorf("expected ')' but got %s at position %d", p.curToken.Type, p.curToken.Pos)
	}

	p.nextToken()
	return expr, nil
}

func (p *Parser) parseIdentifierExpression() (expressions.IExpr, error) {
	identifier := p.curToken.Literal

	if identifier == "beta" {
		p.nextToken()
		return &expressions.IsBeta{}, nil
	}

	p.nextToken()

	if p.curToken.Type == EQUALS {
		p.nextToken()
		if p.curToken.Type != IDENTIFIER {
			return nil, fmt.Errorf("expected value after '==' but got %s at position %d", p.curToken.Type, p.curToken.Pos)
		}

		value := p.curToken.Literal
		p.nextToken()

		switch identifier {
		case "role":
			return &expressions.IsRole{Role: value}, nil
		case "region":
			return &expressions.IsRegion{Region: value}, nil
		default:
			return nil, fmt.Errorf("unknown identifier for equality comparison: %s", identifier)
		}
	} else if p.curToken.Type == LESS_THAN {
		p.nextToken()
		if p.curToken.Type != NUMBER {
			return nil, fmt.Errorf("expected number after '<' but got %s at position %d", p.curToken.Type, p.curToken.Pos)
		}

		numStr := p.curToken.Literal
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", numStr)
		}
		p.nextToken()

		switch identifier {
		case "requests":
			return &expressions.RequestsBelow{Max: num}, nil
		default:
			return nil, fmt.Errorf("unknown identifier for less than comparison: %s", identifier)
		}
	}

	return nil, fmt.Errorf("unexpected token after identifier '%s': %s at position %d", identifier, p.curToken.Type, p.curToken.Pos)
}

func Parse(input string) (expressions.IExpr, error) {
	lexer := NewLexer(input)
	parser := NewParser(lexer)
	return parser.ParseExpression()
}
