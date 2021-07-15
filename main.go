package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/madss/flop-lang/interpreter"
	"github.com/madss/flop-lang/lexer"
	"github.com/madss/flop-lang/parser"
)

func run() error {
	l, err := lexer.New("<stdin>", os.Stdin)
	if err != nil {
		return err
	}

	p, err := parser.New(l)
	if err != nil {
		return err
	}

	decls, err := p.Parse()
	if err != nil {
		return err
	}

	i := interpreter.New()
	return i.Interpret(decls)
}

func main() {
	if err := run(); err != nil {
		var (
			lexerErr  lexer.Error
			parserErr parser.Error
		)
		if ok := errors.As(err, &lexerErr); ok {
			fmt.Fprintf(
				os.Stderr,
				"%s:%d:%d: %s\n",
				lexerErr.Location.Name,
				lexerErr.Location.Line,
				lexerErr.Location.Column,
				lexerErr.Message,
			)
		} else if ok := errors.As(err, &parserErr); ok {
			fmt.Fprintf(
				os.Stderr,
				"%s:%d:%d: %s\n",
				parserErr.Token.Location.Name,
				parserErr.Token.Location.Line,
				parserErr.Token.Location.Column,
				parserErr.Message,
			)
		} else {
			fmt.Fprintf(os.Stderr, err.Error())
		}
	}
}
