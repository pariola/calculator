# calculator

mathematical expression evaluator with the [Shunting Yard Algorithm](https://www.youtube.com/watch?v=Wz85Hiwi5MY).

## Use

```shell
go run main.go
> 5*2*(1*9)
90
> 5+2*(-1*9)+5^2
12
```
To run tests:
```shell
go test ./...
```

### Flow

```text
Input -> Lexical Analysis -> Parsing (Shunting Yard) --> Evaluate --> Output
```