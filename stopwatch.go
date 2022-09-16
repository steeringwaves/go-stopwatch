package stopwatch

import (
	"time"
)

type Stopwatch struct {
	start    time.Time
	timeout  time.Duration
	duration time.Duration
	running  bool
}

// NewStopwatch creates a new stopwatch object
func NewStopwatch() *Stopwatch {
	s := &Stopwatch{time.Time{}, 0, 0, false}
	return s
}

// NewStopwatchWithTimeout creates a new stopwatch object
func NewStopwatchWithTimeout(timeout time.Duration) *Stopwatch {
	s := &Stopwatch{time.Time{}, timeout, 0, false}
	return s
}

// SetTimeout modifies the stopwatch timeout
func (s *Stopwatch) SetTimeout(timeout time.Duration) {
	s.timeout = timeout
}

// Duration returns the elapsed time for the stopwatch only counting
// time periods when the stopwatch is running.
func (s *Stopwatch) Duration() time.Duration {
	d := s.duration

	if true == s.running {
		now := DepGetTime()
		d += now.Sub(s.start)
	}

	return d
}

// Remaining returns the remaining time before the stopwatch timeout has been exceeded.
func (s *Stopwatch) Remaining() time.Duration {
	d := s.Duration()
	r := s.timeout - d

	if r < 0 {
		return 0
	}

	if r > s.timeout {
		return s.timeout
	}

	return r
}

// Expired returns true if the running time has exceeded the timeout.
func (s *Stopwatch) Expired() bool {
	return s.Remaining() <= 0
}

// Stop stops the stopwatch based on the current wall-clock time.
func (s *Stopwatch) Stop() time.Duration {
	s.duration = s.Duration()
	s.running = false

	return s.duration
}

// Start starts the stopwatch based on the current wall-clock time.
func (s *Stopwatch) Start() time.Duration {
	if false == s.running {
		s.running = true
		s.start = DepGetTime()
	}

	return s.duration
}

// Reset clears the stopwatch.
func (s *Stopwatch) Reset() time.Duration {
	s.running = false
	s.duration = 0

	return s.duration
}

// Running returns true if currently running.
func (s *Stopwatch) Running() bool {
	return s.running
}

// Milliseconds returns the elapsed duration in milliseconds.
func (s *Stopwatch) Milliseconds() int64 {
	return int64(s.Duration() * time.Nanosecond / time.Millisecond)
}

// Seconds returns the elapsed duration in seconds.
func (s *Stopwatch) Seconds() int64 {
	return int64(s.Duration() * time.Nanosecond / time.Second)
}

// Minutes returns the elapsed duration in minutes.
func (s *Stopwatch) Minutes() int64 {
	return int64(s.Duration() * time.Nanosecond / time.Minute)
}

// Hours returns the elapsed duration in hours.
func (s *Stopwatch) Hours() int64 {
	return int64(s.Duration() * time.Nanosecond / time.Hour)
}

// Days returns the elapsed duration in days.
func (s *Stopwatch) Days() int64 {
	return int64(s.Duration() * time.Nanosecond / (24 * time.Hour))
}
