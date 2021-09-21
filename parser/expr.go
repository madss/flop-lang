package parser

import (
	"github.com/madss/flop-lang/ast"
	"github.com/madss/flop-lang/token"
)

const (
	lowest = iota
)

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

func (p *Parser) parseExpression(precedense int) (ast.Expression, error) {
	parse := p.prefixes[p.current.Type]
	if parse == nil {
		return nil, p.error("expected expression")
	}
	left, err := parse()
	if err != nil {
		return nil, err
	}
	return left, nil
}

func (p *Parser) parseIdent() (ast.Expression, error) {
	return &ast.IdExpression{Value: p.current}, p.advance()
}

func (p *Parser) parseNum() (ast.Expression, error) {
	return &ast.NumExpression{Value: p.current}, p.advance()
}

func (p *Parser) parseStr() (ast.Expression, error) {
	return &ast.StrExpression{Value: p.current}, p.advance()
}
