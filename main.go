package main

import (
	"fmt"
	"os"

	"calculator/eval"
	"calculator/lexer"
	"calculator/parser"
)

func main() {
	for {
		fmt.Print("> ")

		l := lexer.New(os.Stdin)

		tokens := parser.New(l).Parse()

		fmt.Println(eval.Calculate(tokens))
	}
}
