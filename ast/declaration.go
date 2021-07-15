package ast

import (
	"github.com/madss/flop-lang/token"
)

type Declaration interface {
	isDeclaration()
}

type FnDeclaration struct {
	Name token.Token
	Body []Statement
}

func (d *FnDeclaration) isDeclaration() {}
