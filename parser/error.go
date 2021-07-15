package parser

import "github.com/madss/flop-lang/token"

type Error struct {
	Message string
	Token   token.Token
}

func (e Error) Error() string {
	return e.Message
}
