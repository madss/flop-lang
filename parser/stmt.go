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
	case token.Ident:
		stmt, err = p.parseCallStmt()
	default:
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if err := p.expect(token.Semi, nil, "expected ';'"); err != nil {
		return nil, err
	}
	return stmt, nil
}

func (p *Parser) parseCallStmt() (*ast.CallStatement, error) {
	var call ast.CallStatement

	if err := p.expect(token.Ident, &call.Name, "expected identifier"); err != nil {
		return nil, err
	}

	if err := p.expect(token.LPar, nil, "expected '('"); err != nil {
		return nil, err
	}

	args, err := p.parseExpressionList()
	if err != nil {
		return nil, err
	}
	call.Args = args

	if err := p.expect(token.RPar, nil, "expected ')'"); err != nil {
		return nil, err
	}

	return &call, nil
}
