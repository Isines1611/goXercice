# Intro

`if` statement works like in other popular programming languages but the expression need not be surrounded by parentheses ( ) but the braces { } are required.
There is also `else`, and you can use `else if ... {}`.
Go allows to declare variables in the condition, and use it in the scope:

```go
if v:= add(5,10); v < 100 {
    // v is accessible
}
```

## Futher information

- [if](https://go.dev/tour/flowcontrol/5)
- [If with a short statement](https://go.dev/tour/flowcontrol/6)