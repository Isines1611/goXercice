# Intro

There are pointer in Go but there is no pointer artimetic. 
Pointers in Go only holds the memory address of a value.

```go
var p *int // pointer to an int
i := 42 // set i to 42
p = &i // p holds the memory address of i 
fmt.Println(*p) // read i through the pointer p
*p = 21 // set the value of i through the pointer p
```

## Futher information

- [pointers](https://go.dev/tour/moretypes/1)