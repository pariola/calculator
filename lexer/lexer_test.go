package lexer

import (
	"strings"
	"testing"

	"calculator/token"
)

type testcase struct {
	input  string
	tokens []token.Token
}

func TestLex_Scan(t *testing.T) {

	tests := []testcase{
		{
			input: "2+3",
			tokens: []token.Token{
				{token.NUMBER, "2"},
				{token.PLUS, "+"},
				{token.NUMBER, "3"},
			},
		},
		{
			input: "(2+3)*2",
			tokens: []token.Token{
				{token.O_PARENTHESIS, "("},
				{token.NUMBER, "2"},
				{token.PLUS, "+"},
				{token.NUMBER, "3"},
				{token.C_PARENTHESIS, ")"},
				{token.ASTERISK, "*"},
				{token.NUMBER, "2"},
			},
		},
		{
			input: "2^4-8/2",
			tokens: []token.Token{
				{token.NUMBER, "2"},
				{token.EXPONENT, "^"},
				{token.NUMBER, "4"},
				{token.MINUS, "-"},
				{token.NUMBER, "8"},
				{token.DIVIDE, "/"},
				{token.NUMBER, "2"},
			},
		},
		{
			input: "2.5^4",
			tokens: []token.Token{
				{token.NUMBER, "2.5"},
				{token.EXPONENT, "^"},
				{token.NUMBER, "4"},
			},
		},
		{
			input: "2^4:",
			tokens: []token.Token{
				{token.NUMBER, "2"},
				{token.EXPONENT, "^"},
				{token.NUMBER, "4"},
				{token.ILLEGAL, ":"},
			},
		},
	}

	for i, test := range tests {
		testScan(i, t, test)
	}
}

func testScan(i int, t *testing.T, test testcase) {

	l := New(
		strings.NewReader(test.input),
	)

	var tokens []token.Token

	for {
		tok := l.Scan()

		if tok.Type == token.EOF {
			break
		}

		tokens = append(tokens, tok)
	}

	if len(tokens) != len(test.tokens) {
		t.Errorf("case %d: invalid tokens size, expected %d got %d", i, len(test.tokens), len(tokens))
	}

	for i := 0; i < len(test.tokens); i++ {
		if tokens[i].Type != test.tokens[i].Type {
			t.Errorf("case %d: invalid token type, expected %d got %d", i, test.tokens[i].Type, tokens[i].Type)
		}

		if tokens[i].Literal != test.tokens[i].Literal {
			t.Errorf("case %d: invalid token literal, expected %s got %s", i, test.tokens[i].Literal, tokens[i].Literal)
		}
	}
}
