package ast

import "github.com/madss/flop-lang/token"

type Expression interface {
	isExpression()
}

type IdExpression struct {
	Value token.Token
}

func (e *IdExpression) isExpression() {}

type NumExpression struct {
	Value token.Token
}

func (e *NumExpression) isExpression() {}

type StrExpression struct {
	Value token.Token
}

func (e *StrExpression) isExpression() {}

type BinaryExpression struct {
	Operator    token.Token
	Left, Right Expression
}

func (e *BinaryExpression) isExpression() {}
