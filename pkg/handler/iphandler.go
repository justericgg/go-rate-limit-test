package handler

import (
	"fmt"
	"github.com/justericgg/go-rate-limit-test/pkg/ratelimiter"
	"net"
	"net/http"
	"strconv"
)

func IpHandler(ipLimiter *ratelimiter.IpLimiter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var ip string
		cusIp := r.URL.Query().Get("ip")
		if cusIp != "" {
			ip = cusIp
		} else {
			ip, _, _ = net.SplitHostPort(r.RemoteAddr)
		}

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
