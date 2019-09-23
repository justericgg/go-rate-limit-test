package ratelimiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	mutex         *sync.RWMutex
	tokens        int
	refillNum     int
	lastFillTime  time.Time
	windowTimeSec int
}

func NewTokenBucket(tbNum int, windowTimeSec int) *TokenBucket {
	return &TokenBucket{
		mutex:         &sync.RWMutex{},
		tokens:        tbNum,
		refillNum:     tbNum,
		lastFillTime:  time.Now(),
		windowTimeSec: windowTimeSec,
	}
}

func (tb *TokenBucket) Take(t time.Time) int {

	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	// refill token condition
	if tb.lastFillTime.Add(time.Duration(tb.windowTimeSec) * time.Second).Before(t) {
		tb.tokens = tb.refillNum
		tb.lastFillTime = t
	}

	if tb.tokens == 0 {
		return -1
	}
	tb.tokens = tb.tokens - 1

	return tb.tokens
}
