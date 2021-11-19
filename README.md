# calculator

mathematical expression evaluator with the [Shunting Yard Algorithm](https://www.youtube.com/watch?v=Wz85Hiwi5MY).

This implementation of the Shunting Yard Algorithm has:

```text
- no support for unary operations like '-1' or '1+(-2)'
- no support for functions like sqrt(64)
```

The `calculator` supports the following:

```text
Operations:
- '-' (subtraction)
- '+' (addition)
- '*' (multiplication)
- '/' (division)
- '^' (exponent)

Supports Parenthesis

Number formats:
- integers
- floats
```

## Use

```shell
go run main.go
> 5*2*(1*9)
90
> 5+2*(1*9)+5^2
48 
```

To run tests:

```shell
go test ./...
```

### Flow

```text
Input -> Lexical Analysis -> Parsing (Shunting Yard) --> Evaluate --> Output
```