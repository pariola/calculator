package eval

import (
	"container/list"
	"math"
	"strconv"

	"calculator/token"
)

// Calculate evaluates the tokens Postfix Notation
func Calculate(tokens []token.Token) float64 {

	var stack list.List

	for _, t := range tokens {

		if t.Type == token.NUMBER {
			f, _ := strconv.ParseFloat(t.Literal, 64)
			stack.PushFront(f)
			continue
		}

		var result float64

		b, a := stack.Front(), stack.Front().Next()

		stack.Remove(a)
		stack.Remove(b)

		aValue, bValue := a.Value.(float64), b.Value.(float64)

		switch t.Type {
		case token.PLUS:
			result = aValue + bValue
		case token.MINUS:
			result = aValue - bValue
		case token.ASTERISK:
			result = aValue * bValue
		case token.DIVIDE:
			result = aValue / bValue
		case token.EXPONENT:
			result = math.Pow(aValue, bValue)
		}

		stack.PushFront(result)
	}

	return stack.Front().Value.(float64)
}
