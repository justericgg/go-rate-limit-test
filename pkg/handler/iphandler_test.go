package handler

import (
	"errors"
	"github.com/justericgg/go-rate-limit-test/pkg/ratelimiter"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockResponseWithWriteError struct {
	Code int
}

func (w *mockResponseWithWriteError) Write(buf []byte) (int, error) {
	return 0, errors.New("mock error")
}

func (w *mockResponseWithWriteError) Header() http.Header {
	return http.Header{}
}

func (w *mockResponseWithWriteError) WriteHeader(statusCode int) {
	w.Code = statusCode
}

func TestIpHandler(t *testing.T) {

	t.Run("When request is over limit, response will be Error", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		var ipLimiter = ratelimiter.NewIpLimiter(0, 60)
		IpHandler(ipLimiter)(response, request)

		got := response.Body.String()
		want := "Error"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("When request is over limit, http status code will be 418", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		var ipLimiter = ratelimiter.NewIpLimiter(0, 60)
		IpHandler(ipLimiter)(response, request)

		got := response.Code
		want := http.StatusTeapot

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("When write to response occur error, http status code will be 500", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := mockResponseWithWriteError{}

		var ipLimiter = ratelimiter.NewIpLimiter(0, 60)
		IpHandler(ipLimiter)(&response, request)

		got := response.Code
		want := http.StatusInternalServerError

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Return the count number of token", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		var ipLimiter = ratelimiter.NewIpLimiter(1, 60)
		IpHandler(ipLimiter)(response, request)

		got := response.Body.String()
		want := " 1"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
