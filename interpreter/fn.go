package interpreter

import (
	"github.com/madss/flop-lang/ast"
	"github.com/madss/flop-lang/token"
)

type Fn struct {
	Args []token.Token
	Body []ast.Statement
}
