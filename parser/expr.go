package parser

import (
	"github.com/madss/flop-lang/ast"
	"github.com/madss/flop-lang/token"
)

func (p *Parser) parseExpressionList() ([]ast.Expression, error) {
	var exprs []ast.Expression
	for {
		expr, err := p.parseExpression()
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

func (p *Parser) parseExpression() (ast.Expression, error) {
	switch p.current.Type {
	case token.Str:
		return p.parseString()
	default:
		return nil, nil
	}
}

func (p *Parser) parseString() (*ast.StrExpression, error) {
	var str ast.StrExpression
	if err := p.expect(token.Str, &str.Value, "expected string literal"); err != nil {
		return nil, err
	}
	return &str, nil
}
