package ratelimiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	mutex     *sync.RWMutex
	Tokens    int
	RefillNum int
	Last      time.Time
	LimitSec  int
}

func NewTokenBucket(tbNum int, limitSec int) *TokenBucket {
	return &TokenBucket{
		mutex:     &sync.RWMutex{},
		Tokens:    tbNum,
		RefillNum: tbNum,
		Last:      time.Now(),
		LimitSec:  limitSec,
	}
}

func (tb *TokenBucket) Take(t time.Time) int {

	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	if tb.Last.Add(time.Duration(tb.LimitSec) * time.Second).Before(t) {
		tb.Tokens = tb.RefillNum
	}

	preserveToken := tb.Tokens
	if preserveToken == 0 {
		return preserveToken
	}
	tb.Tokens = tb.Tokens - 1
	tb.Last = t

	return tb.Tokens
}
