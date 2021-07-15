package ast

import "github.com/madss/flop-lang/token"

type Expression interface {
	isExpression()
}

type StrExpression struct {
	Value token.Token
}

func (e *StrExpression) isExpression() {}
