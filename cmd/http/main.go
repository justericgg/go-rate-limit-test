package main

import (
	"github.com/justericgg/go-rate-limit-test/pkg/handler"
	"github.com/justericgg/go-rate-limit-test/pkg/ratelimiter"
	"log"
	"net/http"
)

var limit = 10
var limitSec = 30
var ipLimiter = ratelimiter.NewIpLimiter(limit, limitSec)

func main() {

	http.HandleFunc("/", handler.IpHandler(ipLimiter))

	log.Println("Server listening...")
	log.Fatal(http.ListenAndServe(":80", nil))
}
