package ast

import "github.com/madss/flop-lang/token"

type Expression interface {
	isExpression()
	Token() token.Token
}

type IdExpression struct {
	Value token.Token
}

func (e *IdExpression) isExpression() {}

func (e *IdExpression) Token() token.Token {
	return e.Value
}

type NumExpression struct {
	Value token.Token
}

func (e *NumExpression) isExpression() {}

func (e *NumExpression) Token() token.Token {
	return e.Value
}

type StrExpression struct {
	Value token.Token
}

func (e *StrExpression) isExpression() {}

func (e *StrExpression) Token() token.Token {
	return e.Value
}

type CallExpression struct {
	Fn   Expression
	Args []Expression
}

func (e CallExpression) isExpression() {}

func (e *CallExpression) Token() token.Token {
	return e.Fn.Token()
}

type BinaryExpression struct {
	Operator    token.Token
	Left, Right Expression
}

func (e *BinaryExpression) isExpression() {}

func (e *BinaryExpression) Token() token.Token {
	return e.Operator
}
