# Intro

Go functions can be written to work on multiple types using type parameters. The type parameters of a function appear between brackets, before the function's arguments. 
```go
func Index[T comparable](s []T, x T) int
```

## Futher information

- [Type parameters](https://go.dev/tour/generics/1)
- [Generic types](https://go.dev/tour/generics/2)