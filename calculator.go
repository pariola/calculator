package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"calculator/grammar"
)

type calculatorListener struct {
	*grammar.BaseCalculatorListener
	stack []int
}

func (l *calculatorListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *calculatorListener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *calculatorListener) ExitMulDiv(c *grammar.MulDivContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case grammar.CalculatorLexerDIV:
		l.push(left / right)
	case grammar.CalculatorLexerMUL:
		l.push(left * right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calculatorListener) ExitAddSub(c *grammar.AddSubContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case grammar.CalculatorLexerADD:
		l.push(left + right)
	case grammar.CalculatorLexerSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calculatorListener) ExitPow(_ *grammar.PowContext) {
	right, left := l.pop(), l.pop()

	result := math.Pow(float64(left), float64(right))

	l.push(int(result))
}

func (l *calculatorListener) ExitNumber(c *grammar.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	l.push(i)
}

func main() {
	// is := antlr.NewInputStream("1 * 2 * (3 + 1)")
	is := antlr.NewInputStream("(2 ** 2) + 2")

	// Create the Lexer : can extend lexer to handle errors with Recover()
	lexer := grammar.NewCalculatorLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := grammar.NewCalculatorParser(stream)

	l := &calculatorListener{}
	antlr.ParseTreeWalkerDefault.Walk(l, parser.Start())

	fmt.Printf("=%d\n", l.pop())
}
