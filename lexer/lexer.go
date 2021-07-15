package lexer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/madss/flop-lang/token"
)

type Lexer struct {
	source   *bufio.Reader
	current  rune
	location token.Location
}

func New(name string, source io.Reader) (*Lexer, error) {
	l := Lexer{
		source: bufio.NewReader(source),
		location: token.Location{
			Name:   name,
			Line:   1,
			Column: 1,
		},
	}
	return &l, l.advance()
}

func (l *Lexer) Next() (token.Token, error) {
	t, err := l.lex()
	if errors.Is(err, io.EOF) {
		return l.token(token.Eof, ""), nil
	}
	return t, err
}

func (l *Lexer) lex() (token.Token, error) {
	for {
		switch {
		case unicode.IsSpace(l.current):
			if err := l.advance(); err != nil {
				return token.Token{}, err
			}
			continue
		case l.current == '/':
			if err := l.mustAdvance(); err != nil {
				return token.Token{}, err
			}
			if l.current != '/' {
				return token.Token{}, l.error("unexpected character '%c'. Expected '/'", l.current)
			}
			for l.current != '\n' {
				if err := l.advance(); err != nil {
					return token.Token{}, err
				}
			}
		case l.current == '(':
			return l.advanceAndReturn(token.LPar, "")
		case l.current == ')':
			return l.advanceAndReturn(token.RPar, "")
		case l.current == '{':
			return l.advanceAndReturn(token.LCur, "")
		case l.current == '}':
			return l.advanceAndReturn(token.RCur, "")
		case l.current == ';':
			return l.advanceAndReturn(token.Semi, "")
		case l.current == '"':
			return l.lexString()
		case unicode.IsLetter(l.current) || l.current == '_':
			return l.lexKeywordOrIdentifier()
		default:
			return token.Token{}, l.error("unepected character '%c'", l.current)
		}
	}
}

func (l *Lexer) lexString() (token.Token, error) {
	var str strings.Builder

	for {
		if err := l.mustAdvance(); err != nil {
			return token.Token{}, err
		}
		switch l.current {
		case '"':
			return l.advanceAndReturn(token.Str, str.String())
		case '\\':
			if err := l.mustAdvance(); err != nil {
				return token.Token{}, err
			}
			switch l.current {
			case 'n':
				str.WriteRune('\n')
			default:
				str.WriteRune(l.current)
			}
		default:
			str.WriteRune(l.current)
		}
	}
}

func (l *Lexer) lexKeywordOrIdentifier() (token.Token, error) {
	var str strings.Builder
	for unicode.IsLetter(l.current) || unicode.IsDigit(l.current) || l.current == '_' {
		str.WriteRune(l.current)
		if err := l.advance(); err != nil {
			return token.Token{}, err
		}
	}
	switch str := str.String(); str {
	case "fn":
		return l.token(token.Fn, ""), nil
	default:
		return l.token(token.Ident, str), nil
	}
}

func (l *Lexer) token(t token.Type, val string) token.Token {
	return token.Token{
		Type:     t,
		Value:    val,
		Location: l.location,
	}
}

func (l *Lexer) error(format string, args ...interface{}) error {
	return Error{fmt.Sprintf(format, args...), l.location}
}

func (l *Lexer) advance() error {
	ch, _, err := l.source.ReadRune()
	if err != nil {
		return err
	}
	l.current = ch
	if ch == '\n' {
		l.location.Line++
		l.location.Column = 0
	}
	l.location.Column++
	return nil
}

func (l *Lexer) mustAdvance() error {
	if err := l.advance(); err != nil {
		if errors.Is(err, io.EOF) {
			err = l.error("unexpected end of file")
		}
		return err
	}
	return nil
}

func (l *Lexer) advanceAndReturn(t token.Type, val string) (token.Token, error) {
	return l.token(t, val), l.advance()
}
