package main

import (
	"github.com/justericgg/go-rate-limit-test/pkg/handler"
	"github.com/justericgg/go-rate-limit-test/pkg/ratelimiter"
	"log"
	"net/http"
	"os"
	"strconv"
)

func getEnv(key, def string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return def
	}
	return value
}

var port = getEnv("PORT", "80")
var limit, _ = strconv.Atoi(getEnv("LIMIT", "60"))
var windowTimeSec, _ = strconv.Atoi(getEnv("WINDOW_TIME_SEC", "60"))
var ipLimiter = ratelimiter.NewIpLimiter(limit, windowTimeSec)

func main() {

	http.HandleFunc("/rate-limit", handler.IpHandler(ipLimiter))

	log.Println("Server start at " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
