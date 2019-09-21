package ratelimiter

import (
	"testing"
)

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

	t.Run("Limiter must separate by ip", func(t *testing.T) {

		IpLimiter := NewIpLimiter(1, 60)
		ip1 := "1.1.1.1"
		IpLimiter.Take(ip1)
		IpLimiter.Take(ip1)
		ip2 := "1.1.1.2"

		want := 0
		got := IpLimiter.Take(ip2)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
