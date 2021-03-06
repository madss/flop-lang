package token

type Token struct {
	Type     Type
	Value    string
	Location Location
}

type Type int

const (
	Eof Type = iota
	Ident
	Num
	Str
	Fn
	LPar
	RPar
	LCur
	RCur
	Comma
	Semi
	Plus
	Minus
	Multiply
	Divide
)

type Location struct {
	Name         string
	Line, Column int
}
