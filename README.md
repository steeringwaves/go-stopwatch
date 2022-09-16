# go-stopwatch

![workflow](https://github.com/github/docs/actions/workflows/test.yml/badge.svg)

## usage

```go
import "github.com/steeringwaves/go-stopwatch"

sw := stopwatch.NewStopwatch()
sw.Start()

time.Sleep(250 * time.Millisecond)

fmt.Printf("Stopwatch Duration: %v\n", sw.Duration())

sw.Stop()

time.Sleep(250 * time.Millisecond)
fmt.Printf("Stopwatch Duration: %v\n", sw.Duration())

sw.Start()

time.Sleep(250 * time.Millisecond)
fmt.Printf("Stopwatch Duration: %v\n", sw.Duration())

sw.Reset()
sw.Start()

time.Sleep(250 * time.Millisecond)
fmt.Printf("Stopwatch Duration: %v\n", sw.Duration())
```
