package ratelimiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	mutex    *sync.RWMutex
	Tokens   int
	Last     time.Time
	LimitSec int
}

func NewTokenBucket(tbNum int, limitSec int) *TokenBucket {
	return &TokenBucket{
		mutex:    &sync.RWMutex{},
		Tokens:   tbNum,
		Last:     time.Now(),
		LimitSec: limitSec,
	}
}

func (tb *TokenBucket) Take(t time.Time) int {

	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	preserveToken := tb.Tokens
	if preserveToken == 0 {
		return preserveToken
	}
	tb.Tokens = tb.Tokens - 1
	tb.Last = t

	return preserveToken
}

func (tb *TokenBucket) fill(tbNum int) {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	tb.Tokens = tbNum
	tb.Last = time.Now()
}
