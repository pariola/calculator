package lexer

import (
	"bufio"
	"io"
	"unicode"

	"calculator/token"
)

const eof = byte(0)

type lex struct {
	r *bufio.Reader
}

func New(r io.Reader) *lex {
	return &lex{
		r: bufio.NewReader(r),
	}
}

func (l *lex) unread() {
	_ = l.r.UnreadRune()
}

func (l *lex) read() byte {
	b, err := l.r.ReadByte()

	if err != nil {
		return eof
	}

	return b
}

func (l *lex) readNumber() token.Token {
	panic("unimplemented")
}

func (l lex) Scan() token.Token {

	b := l.read()

	r := rune(b)
	if unicode.IsNumber(r) {
		l.unread()
		return l.readNumber()
	}

	switch b {
	case '+':
		return token.New(token.PLUS, "+")
	case '-':
		return token.New(token.MINUS, "-")
	case '/':
		return token.New(token.DIVIDE, "/")
	case '*':
		return token.New(token.ASTERISK, "*")
	case eof:
		return token.New(token.EOF, "")
	default:
		return token.New(token.ILLEGAL, "")
	}
}
