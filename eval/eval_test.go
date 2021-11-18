package eval

import (
	"testing"

	"calculator/token"
)

func TestCalculate(t *testing.T) {

	// (2+3)*2-3^3+9/3 = -14
	tokens := []token.Token{
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

	if result := Calculate(tokens); result != -14 {
		t.Errorf("invalid result, expected %f got %f", -14.0, result)
	}
}
