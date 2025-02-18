# Intro

mutex allow to make sure only one goroutine can access a variable at a time to avoid conflicts.
Go's standard library provides mutual exclusion with `sync.Mutex` and its two methods: `Lock` and `Unlock`.


## Futher information

- [mutex](https://go.dev/tour/concurrency/9)