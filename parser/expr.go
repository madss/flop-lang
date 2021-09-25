package parser

import (
	"github.com/madss/flop-lang/ast"
	"github.com/madss/flop-lang/token"
)

const (
	lowest = iota
	term
	factor
)

func precedence(t token.Token) int {
	switch t.Type {
	case token.Plus, token.Minus:
		return term
	case token.Multiply, token.Divide:
		return factor
	default:
		return lowest
	}
}

func (p *Parser) parseExpressionList() ([]ast.Expression, error) {
	var exprs []ast.Expression
	for {
		expr, err := p.parseExpression(lowest)
		if err != nil {
			return nil, err
		}
		if expr == nil {
			if len(exprs) > 0 {
				return nil, p.error("expected expression")
			}
			break
		}
		exprs = append(exprs, expr)
		ok, err := p.accept(token.Comma, nil)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
	}
	return exprs, nil
}

func (p *Parser) parseExpression(prec int) (ast.Expression, error) {
	var (
		expr ast.Expression
		err  error
	)
	switch p.current.Type {
	case token.Ident:
		expr, err = &ast.IdExpression{Value: p.current}, p.advance()
	case token.Num:
		expr, err = &ast.NumExpression{Value: p.current}, p.advance()
	case token.Str:
		expr, err = &ast.StrExpression{Value: p.current}, p.advance()
	default:
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	for prec < precedence(p.current) {
		switch p.current.Type {
		case token.Plus, token.Minus, token.Multiply, token.Divide:
			expr, err = p.parseInfix(expr)
			if err != nil {
				return nil, err
			}
		default:
			break
		}
	}
	return expr, nil
}

func (p *Parser) parseInfix(left ast.Expression) (ast.Expression, error) {
	op := p.current
	prec := precedence(op)

	if err := p.advance(); err != nil {
		return nil, err
	}

	right, err := p.parseExpression(prec)
	if err != nil {
		return nil, err
	} else if right == nil {
		return nil, p.error("expected expression")
	}

	return &ast.BinaryExpression{
		Operator: op,
		Left:     left,
		Right:    right,
	}, nil
}
