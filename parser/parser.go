package parser

import (
	"fmt"

	"github.com/madss/flop-lang/ast"
	"github.com/madss/flop-lang/lexer"
	"github.com/madss/flop-lang/token"
)

type Parser struct {
	lexer   *lexer.Lexer
	current token.Token
}

func New(l *lexer.Lexer) (*Parser, error) {
	p := Parser{
		lexer: l,
	}

	return &p, p.advance()
}

func (p *Parser) Parse() ([]ast.Declaration, error) {
	decls, err := p.parseDeclarations()
	if err != nil {
		return nil, err
	}
	err = p.expect(token.Eof, nil, "expected end of file")
	if err != nil {
		return nil, err
	}
	return decls, err
}

func (p *Parser) at(t token.Type) bool {
	return p.current.Type == t
}

func (p *Parser) accept(t token.Type, token *token.Token) (bool, error) {
	accepted := p.at(t)
	if accepted {
		if token != nil {
			*token = p.current
		}
		if err := p.advance(); err != nil {
			return false, err
		}
	}
	return accepted, nil
}

func (p *Parser) expect(t token.Type, token *token.Token, msg string) error {
	ok, err := p.accept(t, token)
	if err != nil {
		return err
	}
	if !ok {
		return p.error(msg)
	}
	return nil
}

func (p *Parser) advance() error {
	t, err := p.lexer.Next()
	if err != nil {
		return err
	}
	p.current = t
	return nil
}

func (p *Parser) error(format string, args ...interface{}) error {
	return Error{fmt.Sprintf(format, args...), p.current}
}
