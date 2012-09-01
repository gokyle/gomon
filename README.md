## monitor

`monitor` provides basic runtime monitoring for Go programs. It is
based on [monitor.py](https://github.com/kisom/pymods/blob/master/monitor.py),
a Python module originally written to ensure a Bitcoin broker continued
running even in the face of unanticipated exceptions. In this case,
it takes a target function with the signature

```go
func target(panicked *bool) error {}
```

When the target function returns, 
