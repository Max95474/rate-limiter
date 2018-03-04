### Go throttle module

Module simply exports one function

Use
```go
fn := limiter.Limit(func() {/*Do stuff*/}, time.Second)
```

to get wrapped function that can be called only once a second.

```bash
go test
```

to run tests