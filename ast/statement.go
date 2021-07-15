package ast

import "github.com/madss/flop-lang/token"

type Statement interface {
	isStatement()
}

type CallStatement struct {
	Name token.Token
	Args []Expression
}

func (s *CallStatement) isStatement() {}
