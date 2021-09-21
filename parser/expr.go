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
	case token.Ident:
		return &ast.IdExpression{Value: p.current}, p.advance()
	case token.Num:
		return &ast.NumExpression{Value: p.current}, p.advance()
	case token.Str:
		return &ast.StrExpression{Value: p.current}, p.advance()
	default:
		return nil, nil
	}
}
