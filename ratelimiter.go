package ratelimiter

import (
	"time"
)

type throttle struct {
	period time.Duration
	last time.Time
}

// Returns the same function but wrapped into
// throttle function that can be called only
// within allowed period of time
func Limit(fn func(), period time.Duration) func() {
	throttle := throttle{
		period: period,
		last: time.Now().Add(-period),
	}

	return func() {
		allowedTime := throttle.last.Add(throttle.period)
		if time.Now().After(allowedTime) {
			fn()
			throttle.last = time.Now()
		}
	}
}
