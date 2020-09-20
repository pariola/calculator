// Code generated from grammar/Calculator.g4 by ANTLR 4.8. DO NOT EDIT.

package grammar

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 48, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 6, 9,
	38, 10, 9, 13, 9, 14, 9, 39, 3, 10, 6, 10, 43, 10, 10, 13, 10, 14, 10,
	44, 3, 10, 3, 10, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9,
	17, 10, 19, 11, 3, 2, 4, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2,
	49, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2,
	2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2,
	2, 2, 2, 19, 3, 2, 2, 2, 3, 21, 3, 2, 2, 2, 5, 23, 3, 2, 2, 2, 7, 25, 3,
	2, 2, 2, 9, 27, 3, 2, 2, 2, 11, 29, 3, 2, 2, 2, 13, 31, 3, 2, 2, 2, 15,
	33, 3, 2, 2, 2, 17, 37, 3, 2, 2, 2, 19, 42, 3, 2, 2, 2, 21, 22, 7, 42,
	2, 2, 22, 4, 3, 2, 2, 2, 23, 24, 7, 43, 2, 2, 24, 6, 3, 2, 2, 2, 25, 26,
	7, 44, 2, 2, 26, 8, 3, 2, 2, 2, 27, 28, 7, 49, 2, 2, 28, 10, 3, 2, 2, 2,
	29, 30, 7, 45, 2, 2, 30, 12, 3, 2, 2, 2, 31, 32, 7, 47, 2, 2, 32, 14, 3,
	2, 2, 2, 33, 34, 7, 44, 2, 2, 34, 35, 7, 44, 2, 2, 35, 16, 3, 2, 2, 2,
	36, 38, 9, 2, 2, 2, 37, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 37, 3,
	2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 18, 3, 2, 2, 2, 41, 43, 9, 3, 2, 2, 42,
	41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2,
	2, 45, 46, 3, 2, 2, 2, 46, 47, 8, 10, 2, 2, 47, 20, 3, 2, 2, 2, 5, 2, 39,
	44, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "'**'",
}

var lexerSymbolicNames = []string{
	"", "", "", "MUL", "DIV", "ADD", "SUB", "POW", "NUMBER", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "MUL", "DIV", "ADD", "SUB", "POW", "NUMBER", "WHITESPACE",
}

type CalculatorLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCalculatorLexer(input antlr.CharStream) *CalculatorLexer {

	l := new(CalculatorLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Calculator.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalculatorLexer tokens.
const (
	CalculatorLexerT__0       = 1
	CalculatorLexerT__1       = 2
	CalculatorLexerMUL        = 3
	CalculatorLexerDIV        = 4
	CalculatorLexerADD        = 5
	CalculatorLexerSUB        = 6
	CalculatorLexerPOW        = 7
	CalculatorLexerNUMBER     = 8
	CalculatorLexerWHITESPACE = 9
)
