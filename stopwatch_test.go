package stopwatch

import (
	"fmt"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

func TestStopwatchNew(t *testing.T) {

	g := Goblin(t)

	g.Describe("stopwatch.NewStopwatch()", func() {

		g.It("Should create new stopwatch", func(done Done) {
			sw := NewStopwatch()

			g.Assert(sw != nil).Equal(true)
			g.Assert(time.Duration(0)).Equal(sw.duration)
			g.Assert(false).Equal(sw.running)
			g.Assert(true).Equal(sw.start.IsZero())

			done()
		})
	})
}

var verifyDurationTests = []struct {
	it       string
	increase time.Duration
	expected time.Duration
	stop     bool
	reset    bool
}{
	{
		it:       "Should increase by 1 minute",
		increase: 1 * time.Minute,
		expected: (1 * time.Minute) * time.Nanosecond,
		stop:     false,
		reset:    false,
	},
	{
		it:       "Should increase by 1 minute",
		increase: 1 * time.Minute,
		expected: (2 * time.Minute) * time.Nanosecond,
		stop:     false,
		reset:    false,
	},
	{
		it:       "Should pause",
		increase: 1 * time.Minute,
		expected: (2 * time.Minute) * time.Nanosecond,
		stop:     true,
		reset:    false,
	},
	{
		it:       "Should increase by 1 minute",
		increase: 1 * time.Minute,
		expected: (3 * time.Minute) * time.Nanosecond,
		stop:     false,
		reset:    false,
	},
	{
		it:       "Should increase by 1 minute",
		increase: 1 * time.Minute,
		expected: (4 * time.Minute) * time.Nanosecond,
		stop:     false,
		reset:    false,
	},
	{
		it:       "Should reset to zero",
		increase: 0 * time.Minute,
		expected: (0 * time.Minute) * time.Nanosecond,
		stop:     false,
		reset:    true,
	},
	{
		it:       "Should increase by 1 minute",
		increase: 1 * time.Minute,
		expected: (1 * time.Minute) * time.Nanosecond,
		stop:     false,
		reset:    false,
	},
}

func TestStopwatch(t *testing.T) {
	g := Goblin(t)

	Stubs()
	sw := NewStopwatch()
	sw.Start()

	g.Describe("stopwatch.Duration() using fake clock", func() {
		for _, tt := range verifyDurationTests {

			testdata := tt

			g.It(testdata.it, func(done Done) {
				if true == testdata.reset {
					sw.Reset()
				}

				if true == testdata.stop {
					sw.Stop()
				} else {
					sw.Start()
				}

				_stubClockAdvance(testdata.increase)

				result := sw.Duration()

				g.Assert(result).Equal(testdata.expected)

				ms := int64((testdata.expected / time.Millisecond))
				g.Assert(ms).Equal(sw.Milliseconds())

				secs := int64((testdata.expected / time.Second))
				g.Assert(secs).Equal(sw.Seconds())

				mins := int64((testdata.expected / time.Minute))
				g.Assert(mins).Equal(sw.Minutes())

				hours := int64((testdata.expected / time.Hour))
				g.Assert(hours).Equal(sw.Hours())

				days := int64((testdata.expected / (24 * time.Hour)))
				g.Assert(days).Equal(sw.Days())

				done()
			})
		}
	})

	StubsRestore()
}

func TestStopwatch_Real(t *testing.T) {
	g := Goblin(t)

	g.Describe("stopwatch.Duration() using fake clock", func() {

		g.It("Should increase", func(done Done) {
			sw := NewStopwatch()
			sw.Start()

			time.Sleep(1 * time.Millisecond)

			result := int64(sw.Duration())
			g.Assert(result > 0).Equal(true)
			done()
		})
	})
}

func ExampleStopwatch() {
	sw := NewStopwatch()
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
}
