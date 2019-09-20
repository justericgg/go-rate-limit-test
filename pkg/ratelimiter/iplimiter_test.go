package ratelimiter

import "testing"

func TestIpLimiterTake(t *testing.T) {

	t.Run("A new ip have to crate a new bucket and get token", func(t *testing.T) {

		ipLimiter := NewIpLimiter(60, 60)
		ip := "8.8.8.8"
		token := ipLimiter.Take(ip)

		if token != 59 {
			t.Errorf("got %v want %v", token, 59)
		}

		if _, ok := ipLimiter.ips[ip]; !ok {
			t.Errorf("ip must set in map")
		}

	})

}
