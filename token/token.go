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
	Str
	Fn
	LPar
	RPar
	LCur
	RCur
	Comma
	Semi
)

type Location struct {
	Name         string
	Line, Column int
}
