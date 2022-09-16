package stopwatch

import (
	"time"

	"github.com/jonboulle/clockwork"
)

var oldDepGetTime = DepGetTime
var _stubClock clockwork.FakeClock
var _stubClockAdvance = func(t time.Duration) {}

func Stubs() {
	_stubClock = clockwork.NewFakeClock()

	DepGetTime = func() time.Time {
		return _stubClock.Now()
	}

	_stubClockAdvance = func(t time.Duration) {
		_stubClock.Advance(t)
	}
}

func StubsRestore() {
	DepGetTime = oldDepGetTime
}
