package ratelimiter

import (
	"testing"
	"time"
)

func TestLimitFunc(t *testing.T) {
	var calledTimes = 0
	fn := Limit(func() {
		calledTimes++
	}, time.Second, 0)
	fn()
	if calledTimes > 1 {
		t.Fatal("Function called twice!")
	}
	time.Sleep(time.Millisecond * 300)
	fn()
	if calledTimes > 1 {
		t.Fatal("Function called more than one time!")
	}
	time.Sleep(time.Second)
	fn()
	if calledTimes != 2 {
		t.Fatal("Function isn't called second time!")
	}
}

func TestCalledTimes(t *testing.T) {
	var calledTimesFirts, calledTimesSecond int
	fn1 := Limit(func() {
		calledTimesFirts++
	}, time.Second, 5)
	for i := 0; i < 10; i++ {
		fn1()
	}
	if calledTimesFirts > 5 {
		t.Fatal("Function called more than 5 times")
	}
	time.Sleep(time.Second)
	fn2 := Limit(func() {
		calledTimesSecond++
	}, time.Second, 3)
	for i := 0; i < 10; i++ {
		fn2()
	}
	if calledTimesSecond > 3 {
		t.Fatal("Function called more than 5 times")
	}
}