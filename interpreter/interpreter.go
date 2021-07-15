package interpreter

import (
	"fmt"

	"github.com/madss/flop-lang/ast"
	"github.com/madss/flop-lang/interpreter/env"
	"github.com/madss/flop-lang/token"
)

type Interpreter struct {
	env *env.Environment
}

func New() *Interpreter {
	env := env.New()
	env.Set("print", Builtin(print))
	return &Interpreter{env}
}

func (i *Interpreter) Interpret(decls []ast.Declaration) error {
	if err := i.interpretDeclarations(i.env, decls); err != nil {
		return err
	}
	callMain := ast.CallStatement{
		Name: token.Token{Type: token.Ident, Value: "main"},
	}
	if err := i.interpretStatement(i.env, &callMain); err != nil {
		return err
	}
	return nil
}

func (i *Interpreter) interpretDeclarations(env *env.Environment, decls []ast.Declaration) error {
	for _, decl := range decls {
		switch decl := decl.(type) {
		case *ast.FnDeclaration:
			i.env.Set(decl.Name.Value, Fn{decl.Body})
		default:
			panic("unepxected declaration")
		}
	}
	return nil
}

func (i *Interpreter) interpretStatements(env *env.Environment, stmts []ast.Statement) error {
	for _, stmt := range stmts {
		if err := i.interpretStatement(env, stmt); err != nil {
			return err
		}
	}
	return nil
}

func (i *Interpreter) interpretStatement(env *env.Environment, stmt ast.Statement) error {
	switch stmt := stmt.(type) {
	case *ast.CallStatement:
		fn := env.Get(stmt.Name.Value)
		if fn == nil {
			return i.error(stmt.Name, "couldn't find function %s", stmt.Name.Value)
		}

		var args []interface{}
		for _, arg := range stmt.Args {
			arg, err := i.interpretExpression(env, arg)
			if err != nil {
				return err
			}
			args = append(args, arg)
		}

		switch fn := fn.(type) {
		case Builtin:
			return fn(args...)
		case Fn:
			childEnv := env.Child()
			return i.interpretStatements(childEnv, fn.Body)
		default:
			return i.error(stmt.Name, "%s is not a function", stmt.Name.Value)
		}
	default:
		panic("unexpected statement")
	}
	return nil
}

func (i *Interpreter) interpretExpression(env *env.Environment, expr ast.Expression) (interface{}, error) {
	switch expr := expr.(type) {
	case *ast.StrExpression:
		return expr.Value.Value, nil
	default:
		panic("unexpected expression")
	}
}

func (i *Interpreter) error(token token.Token, format string, args ...interface{}) error {
	return fmt.Errorf("%s (%s:%s)", fmt.Sprintf(format, args...), token.Location.Name, token.Location.Line)
}
