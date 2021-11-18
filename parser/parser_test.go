package parser

import (
	"testing"

	"calculator/token"
)

// lexer
type lexer struct {
	tokens []token.Token
}

func (l *lexer) Scan() token.Token {

	if len(l.tokens) == 0 {
		return token.New(token.EOF, "")
	}

	head := l.tokens[0]
	l.tokens = l.tokens[1:]

	return head
}

func TestParser_Parse(t *testing.T) {

	expected := []token.Token{
		{token.NUMBER, "2"},
		{token.NUMBER, "3"},
		{token.PLUS, "+"},
		{token.NUMBER, "2"},
		{token.ASTERISK, "*"},
		{token.NUMBER, "3"},
		{token.NUMBER, "3"},
		{token.EXPONENT, "^"},
		{token.MINUS, "-"},
		{token.NUMBER, "9"},
		{token.NUMBER, "3"},
		{token.DIVIDE, "/"},
		{token.PLUS, "+"},
	}

	// (2+3)*2-3^3+9/3
	l := &lexer{
		tokens: []token.Token{
			{token.O_PARENTHESIS, "("},
			{token.NUMBER, "2"},
			{token.PLUS, "+"},
			{token.NUMBER, "3"},
			{token.C_PARENTHESIS, ")"},
			{token.ASTERISK, "*"},
			{token.NUMBER, "2"},
			{token.MINUS, "-"},
			{token.NUMBER, "3"},
			{token.EXPONENT, "^"},
			{token.NUMBER, "3"},
			{token.PLUS, "+"},
			{token.NUMBER, "9"},
			{token.DIVIDE, "/"},
			{token.NUMBER, "3"},
		},
	}

	tokens := New(l).Parse()

	if len(tokens) != len(expected) {
		t.Errorf("invalid tokens size, expected %d got %d", len(expected), len(tokens))
	}

	for i := 0; i < len(expected); i++ {
		if tokens[i].Type != expected[i].Type {
			t.Errorf("token %d: invalid token type, expected %d got %d", i, expected[i].Type, tokens[i].Type)
		}

		if tokens[i].Literal != expected[i].Literal {
			t.Errorf("token %d: invalid token literal, expected %s got %s", i, expected[i].Literal, tokens[i].Literal)
		}
	}
}
