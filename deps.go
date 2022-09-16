package stopwatch

import (
	"time"
)

// DepGetTime is an external dependancy wrapper
// that returns a time.Time value using
// the time.Now() function
var DepGetTime = func() time.Time {
	return (time.Now())
}
