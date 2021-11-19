package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"calculator/eval"
	"calculator/lexer"
	"calculator/parser"
)

func main() {
	for {
		fmt.Print("> ")

		r := bufio.NewReader(os.Stdin)
		in, _ := r.ReadString('\n')

		l := lexer.New(
			strings.NewReader(in),
		)

		tokens := parser.New(l).Parse()

		fmt.Println(eval.Calculate(tokens))
	}
}
