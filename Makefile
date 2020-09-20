build-ast:
	antlr -Dlanguage=Go -package grammar grammar/Calculator.g4
