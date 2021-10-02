package parser

import (
	"github.com/madss/flop-lang/ast"
	"github.com/madss/flop-lang/token"
)

func (p *Parser) parseStatements() ([]ast.Statement, error) {
	var stmts []ast.Statement
	for {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		if stmt == nil {
			break
		}
		stmts = append(stmts, stmt)
	}
	return stmts, nil
}

func (p *Parser) parseStatement() (ast.Statement, error) {
	var (
		stmt ast.Statement
		err  error
	)
	switch p.current.Type {
	default:
		expr, err := p.parseExpression(lowest)
		if err != nil {
			return nil, err
		}
		if expr == nil {
			return nil, nil
		}
		stmt = &ast.ExpressionStatement{Expr: expr}
	}
	if err != nil {
		return nil, err
	}
	if err := p.expect(token.Semi, nil, "expected ';'"); err != nil {
		return nil, err
	}
	return stmt, nil
}
