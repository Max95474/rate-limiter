### Simple Go throttle module

Module simply exports one function

Use
```go
fn := limiter.Limit(func() {/*Do stuff*/}, time.Second, 10)
```

to get wrapped function that can be called only once a second.

The Limit function accepts:
* function of type func()
* time.Duration that specifies period of time
* max calls that can be performed within specified period

```bash
go test
```

to run tests