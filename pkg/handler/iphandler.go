package handler

import (
	"fmt"
	"github.com/justericgg/go-rate-limit-test/pkg/iptool"
	"github.com/justericgg/go-rate-limit-test/pkg/ratelimiter"
	"net/http"
	"strconv"
)

func IpHandler(ipLimiter *ratelimiter.IpLimiter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := iptool.ClientIP(r)
		token := ipLimiter.Take(ip)

		if token == -1 {
			w.WriteHeader(http.StatusTeapot)
			_, err := w.Write([]byte("Error"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		msg := ip + " " + strconv.Itoa(ipLimiter.GetLimit()-token)
		_, _ = fmt.Fprintf(w, msg)
		return
	}
}
