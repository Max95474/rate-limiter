package ratelimiter

import (
	"time"
)

type throttle struct {
	period time.Duration
	maxCalls int
	last time.Time
	callsCount int
}

// Returns the same function but wrapped into
// throttle function that will be called only
// within allowed period of time
// Otherwise the function call will be ignored
func Limit(fn func(), period time.Duration, maxCalls int) func() {
	throttle := throttle{
		period: period,
		maxCalls: maxCalls,
		last: time.Now().Add(-period),
	}

	return func() {
		allowedTime := throttle.last.Add(throttle.period)
		if time.Now().After(allowedTime) && throttle.callsCount < throttle.maxCalls {
			fn()
			throttle.last = time.Now()
			throttle.maxCalls++
		}
	}
}
