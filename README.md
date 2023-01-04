# go-stopwatch
![workflow](https://github.com/steeringwaves/go-stopwatch/actions/workflows/test.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/steeringwaves/go-stopwatch.svg)](https://pkg.go.dev/github.com/steeringwaves/go-stopwatch)

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
