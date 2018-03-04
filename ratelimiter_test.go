package ratelimiter

import (
	"testing"
	"time"
	"fmt"
)

func TestLimitFunc(t *testing.T) {
	var calledTimes = 0
	fn := Limit(func() {
		calledTimes = calledTimes + 1
		fmt.Println("Running code...", calledTimes)
	}, time.Second)
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