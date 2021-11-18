package parser

import (
	"container/list"

	"calculator/token"
)

type Lexer interface {
	Scan() token.Token
}

var (
	pemdas = map[token.Type]int{
		token.O_PARENTHESIS: 5,
		token.EXPONENT:      4,
		token.ASTERISK:      3,
		token.DIVIDE:        3,
		token.PLUS:          2,
		token.MINUS:         2,
	}
)

// Parser implements the Shunting Yard Algorithm
type Parser struct {
	l Lexer

	// s the operator stack, head is at front
	s list.List

	// q the numbers queue
	q list.List
}

// New returns instance of Parser
func New(l Lexer) *Parser {
	return &Parser{l: l}
}

// Parse convert Infix tokens from Lexer to Postfix Notation
func (p *Parser) Parse() []token.Token {

	for {
		t := p.l.Scan()

		if t.Type == token.EOF || t.Type == token.ILLEGAL {
			break
		}

		p.parseToken(t)
	}

	// pop and enqueue remaining operators
	op := p.s.Front()
	for op != nil {
		p.q.PushBack(op.Value)
		op = op.Next()
	}

	var tokens []token.Token

	// read
	for f := p.q.Front(); f != nil; f = f.Next() {
		tokens = append(tokens, f.Value.(token.Token))
	}

	return tokens
}

// parseToken
func (p *Parser) parseToken(t token.Token) {

	switch t.Type {
	case token.NUMBER:
		p.q.PushBack(t)
		return
	case token.C_PARENTHESIS:
		for op := p.s.Front(); op != nil; op = p.s.Front() {

			p.s.Remove(op)

			if op.Value.(token.Token).Type == token.O_PARENTHESIS {
				break
			}

			// enqueue popped operators
			p.q.PushBack(op.Value)
		}
		return
	}

	// an operator
	for elem := p.s.Front(); elem != nil; elem = p.s.Front() {

		head := elem.Value.(token.Token)

		// compare precedence with PEMDAS
		if head.Type == token.O_PARENTHESIS || pemdas[head.Type] < pemdas[t.Type] {
			break
		}

		// pop stack and enqueue
		p.s.Remove(elem)
		p.q.PushBack(head)
	}

	// push operator to stack
	p.s.PushFront(t)
}
