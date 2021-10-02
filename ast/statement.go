package ast

type Statement interface {
	isStatement()
}

type ExpressionStatement struct {
	Expr Expression
}

func (s *ExpressionStatement) isStatement() {}
