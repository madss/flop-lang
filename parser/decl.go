package parser

import (
	"github.com/madss/flop-lang/ast"
	"github.com/madss/flop-lang/token"
)

func (p *Parser) parseDeclarations() ([]ast.Declaration, error) {
	var decls []ast.Declaration
	for {
		decl, err := p.parseDeclaration()
		if err != nil {
			return nil, err
		}
		if decl == nil {
			break
		}
		decls = append(decls, decl)
	}
	return decls, nil
}

func (p *Parser) parseDeclaration() (ast.Declaration, error) {
	switch {
	case p.at(token.Fn):
		return p.parseFunction()
	default:
		return nil, nil
	}
}

func (p *Parser) parseFunction() (*ast.FnDeclaration, error) {
	fn := ast.FnDeclaration{}

	if err := p.expect(token.Fn, nil, "expected 'fn' keyword"); err != nil {
		return nil, err
	}

	if err := p.expect(token.Ident, &fn.Name, "expected identifier"); err != nil {
		return nil, err
	}

	if err := p.expect(token.LPar, nil, "expected '('"); err != nil {
		return nil, err
	}

	args, err := p.parseIdentifierList()
	if err != nil {
		return nil, err
	}
	fn.Args = args

	if err := p.expect(token.RPar, nil, "expected ')'"); err != nil {
		return nil, err
	}

	if err := p.expect(token.LCur, nil, "expected '{'"); err != nil {
		return nil, err
	}

	stmts, err := p.parseStatements()
	if err != nil {
		return nil, err
	}
	fn.Body = stmts

	if err := p.expect(token.RCur, nil, "expected '}'"); err != nil {
		return nil, err
	}

	return &fn, nil
}

func (p *Parser) parseIdentifierList() ([]token.Token, error) {
	var ids []token.Token
	for {
		var id token.Token
		ok, err := p.accept(token.Ident, &id)
		if err != nil {
			return nil, err
		}
		if !ok {
			if len(ids) > 0 {
				return nil, p.error("expected identifier")
			}
			break
		}
		ids = append(ids, id)
		ok, err = p.accept(token.Comma, nil)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
	}
	return ids, nil
}
