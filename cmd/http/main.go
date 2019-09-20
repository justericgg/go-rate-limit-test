package main

import (
	"github.com/justericgg/go-rate-limit-test/pkg/handler"
	"github.com/justericgg/go-rate-limit-test/pkg/ratelimiter"
	"log"
	"net/http"
	"os"
)

var limit = 10
var limitSec = 30
var ipLimiter = ratelimiter.NewIpLimiter(limit, limitSec)

func main() {

	http.HandleFunc("/rate-limit", handler.IpHandler(ipLimiter))

	log.Println("Server listening...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
