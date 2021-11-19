package lexer

import (
	"bufio"
	"bytes"
	"io"
	"unicode"

	"calculator/token"
)

const eof = rune(0)

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

func (l *lex) read() rune {
	b, _, err := l.r.ReadRune()

	if err != nil {
		return eof
	}

	return b
}

func (l *lex) skipWhitespace() {
	for {
		r := l.read()

		if !unicode.IsSpace(r) {
			l.unread()
			break
		}
	}
}

func (l *lex) readNumber() token.Token {

	var floating bool
	var buf bytes.Buffer

	for {
		r := l.read()

		if r == '.' && !floating {
			floating = true
		} else if !unicode.IsDigit(r) {
			l.unread()
			break
		}

		buf.WriteRune(r)
	}

	return token.New(token.NUMBER, buf.String())
}

func (l *lex) Scan() token.Token {

	// skip whitespaces early
	l.skipWhitespace()

	r := l.read()

	switch r {
	case '+':
		return token.New(token.PLUS, "+")
	case '-':
		return token.New(token.MINUS, "-")
	case '/':
		return token.New(token.DIVIDE, "/")
	case '^':
		return token.New(token.EXPONENT, "^")
	case '*':
		return token.New(token.ASTERISK, "*")
	case '(':
		return token.New(token.O_PARENTHESIS, "(")
	case ')':
		return token.New(token.C_PARENTHESIS, ")")
	case eof:
		return token.New(token.EOF, "")
	default:

		if unicode.IsNumber(r) {
			l.unread()
			return l.readNumber()
		}

		return token.New(token.ILLEGAL, string(r))
	}
}
