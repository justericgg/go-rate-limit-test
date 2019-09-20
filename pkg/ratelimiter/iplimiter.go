package ratelimiter

import (
	"sync"
	"time"
)

type IpLimiter struct {
	mutex    *sync.RWMutex
	ips      map[string]*TokenBucket
	limit    int
	limitSec int
}

func NewIpLimiter(limit int, limitSec int) *IpLimiter {
	return &IpLimiter{
		mutex:    &sync.RWMutex{},
		ips:      make(map[string]*TokenBucket),
		limit:    limit,
		limitSec: limitSec,
	}
}

func (ipL *IpLimiter) Take(ip string) int {

	if _, ok := ipL.ips[ip]; !ok {
		tb := NewTokenBucket(ipL.limit, ipL.limitSec)
		ipL.ips[ip] = tb
	}

	token := ipL.ips[ip].Take(time.Now())

	return token
}

func (ipL *IpLimiter) GetLimit() int {
	return ipL.limit
}
