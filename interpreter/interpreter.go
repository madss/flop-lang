package interpreter

import (
	"fmt"
	"strconv"

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
	callMain := ast.ExpressionStatement{
		Expr: &ast.CallExpression{
			Fn: &ast.IdExpression{
				token.Token{Type: token.Ident, Value: "main"},
			},
		},
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
			i.env.Set(decl.Name.Value, Fn{decl.Args, decl.Body})
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
	case *ast.ExpressionStatement:
		_, err := i.interpretExpression(env, stmt.Expr)
		if err != nil {
			return err
		}
	default:
		panic("unexpected statement")
	}
	return nil
}

func (i *Interpreter) interpretExpression(env *env.Environment, expr ast.Expression) (interface{}, error) {
	switch expr := expr.(type) {
	case *ast.IdExpression:
		return env.Get(expr.Value.Value), nil
	case *ast.NumExpression:
		val, err := strconv.Atoi(expr.Value.Value)
		if err != nil {
			panic(err)
		}
		return val, nil
	case *ast.StrExpression:
		return expr.Value.Value, nil
	case *ast.CallExpression:
		fn, err := i.interpretExpression(env, expr.Fn)
		if err != nil {
			return nil, err
		}

		var args []interface{}
		for _, arg := range expr.Args {
			arg, err := i.interpretExpression(env, arg)
			if err != nil {
				return nil, err
			}
			args = append(args, arg)
		}

		switch fn := fn.(type) {
		case Builtin:
			return fn(args...), nil
		case Fn:
			childEnv := env.Child()
			if len(fn.Args) != len(args) {
				return nil, i.error(expr.Fn.Token(), "expected %d arguments", len(args))
			}
			for i := range args {
				childEnv.Set(fn.Args[i].Value, args[i])
			}
			return nil, i.interpretStatements(childEnv, fn.Body)
		default:
			return nil, i.error(expr.Fn.Token(), "Calling non-function")
		}
	case *ast.BinaryExpression:
		left, err := i.interpretExpression(env, expr.Left)
		if err != nil {
			return nil, err
		}
		right, err := i.interpretExpression(env, expr.Right)
		if err != nil {
			return nil, err
		}
		leftVal, leftOk := left.(int)
		rightVal, rightOk := right.(int)
		if !leftOk || !rightOk {
			return nil, i.error(expr.Operator, "performing binary operation on non-integer")
		}
		switch expr.Operator.Type {
		case token.Plus:
			return leftVal + rightVal, nil
		case token.Minus:
			return leftVal - rightVal, nil
		case token.Multiply:
			return leftVal * rightVal, nil
		case token.Divide:
			if rightVal == 0 {
				return nil, i.error(expr.Operator, "dividing with zero")
			}
			return leftVal / rightVal, nil
		default:
			panic("unknown binary operator")
		}
	default:
		panic("unexpected expression")
	}
}

func (i *Interpreter) error(token token.Token, format string, args ...interface{}) error {
	return fmt.Errorf("%s (%s:%d)", fmt.Sprintf(format, args...), token.Location.Name, token.Location.Line)
}
