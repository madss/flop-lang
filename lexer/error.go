package lexer

import "github.com/madss/flop-lang/token"

type Error struct {
	Message  string
	Location token.Location
}

func (e Error) Error() string {
	return e.Message
}
