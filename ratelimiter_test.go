package ratelimiter

import (
	"testing"
	"time"
)

func TestLimitFunc(t *testing.T) {
	var calledTimes = 0
	fn := Limit(func() {
		calledTimes++
	}, time.Second, 10)
	fn()
	if calledTimes > 1 {
		t.Fatal("Function called twice!")
	}
	time.Sleep(time.Millisecond * 500)
	fn()
	if calledTimes > 1 {
		t.Fatal("Function called twice!")
	}
	time.Sleep(time.Second)
	fn()
	if calledTimes != 2 {
		t.Fatal("Function isn't called second time!")
	}
}

func TestCalledTimes(t *testing.T) {
	var calledTimes = 0
	fn := Limit(func() {
		calledTimes++
	}, time.Second, 5)
	for i := 0; i < 10; i++ {
		fn()
	}
	if calledTimes > 5 {
		t.Fatal("Function called more than 5 times")
	}
}