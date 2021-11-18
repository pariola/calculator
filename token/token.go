package token

type Type int

const (
	ILLEGAL Type = iota
	EOF

	// Operators
	PLUS     // +
	MINUS    // -
	DIVIDE   // /
	ASTERISK // *

	// Numbers -> Integers | Float
	NUMBER // 12.32, 1560
)

type Token struct {
	Type    Type
	Literal string
}

func New(t Type, l string) Token {
	return Token{Type: t, Literal: l}
}
