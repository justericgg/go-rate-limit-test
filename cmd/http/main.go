package main

import (
	"fmt"
	"github.com/justericgg/go-rate-limit-test/pkg/ratelimiter"
	"log"
	"net/http"
	"strconv"
	"time"
)

var tokenBucket = ratelimiter.NewTokenBucket(60, 60)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		token := tokenBucket.Take(time.Now())
		log.Println(token)
		_, _ = fmt.Fprintf(w, strconv.Itoa(token))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
