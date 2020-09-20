// Code generated from grammar/Calculator.g4 by ANTLR 4.8. DO NOT EDIT.

package grammar // Calculator
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalculatorListener is a complete listener for a parse tree produced by CalculatorParser.
type CalculatorListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterPow is called when entering the Pow production.
	EnterPow(c *PowContext)

	// EnterParenthesis is called when entering the parenthesis production.
	EnterParenthesis(c *ParenthesisContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitPow is called when exiting the Pow production.
	ExitPow(c *PowContext)

	// ExitParenthesis is called when exiting the parenthesis production.
	ExitParenthesis(c *ParenthesisContext)
}
