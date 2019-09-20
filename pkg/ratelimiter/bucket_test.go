package ratelimiter

import (
	"testing"
	"time"
)

func TestNewTokenBucket(t *testing.T) {
	type args struct {
		tbNum    int
		limitSec int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Basic NewTokenBucket",
			args: args{60, 60},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tb := NewTokenBucket(tt.args.tbNum, tt.args.limitSec)
			if tb.Tokens != tt.args.tbNum {
				t.Errorf("got %v, want %v", tb.Tokens, tt.args.tbNum)
			}
			if tb.LimitSec != tt.args.limitSec {
				t.Errorf("got %v, want %v", tb.LimitSec, tt.args.limitSec)
			}
		})
	}
}

func TestBucketTake(t *testing.T) {
	t.Run("When token is zero must return -1 and do not minus one", func(t *testing.T) {
		tb := NewTokenBucket(0, 60)
		takeTime, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 00:00:00")
		r := tb.Take(takeTime)
		expected := -1

		if r != -1 {
			t.Errorf("got %v, want %v", r, expected)
		}

		if tb.Tokens < 0 {
			t.Errorf("got %v, want %v", tb.Tokens, expected)
		}
	})

	t.Run("Taking one token and check tokens will decrease one token", func(t *testing.T) {
		tb := NewTokenBucket(1, 60)
		takeTime, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 00:00:00")
		r := tb.Take(takeTime)
		expected := 0
		if r != 0 {
			t.Errorf("got %v, want %v", r, expected)
		}
		if tb.Tokens != 0 {
			t.Errorf("got %v, want %v", tb.Tokens, expected)
		}
		if tb.Last != takeTime {
			t.Errorf("got %v, want %v", tb.Last, expected)
		}
	})

	t.Run("When none of tokens in the bucket but now is after 1 min from last taking time must fill the bucket", func(t *testing.T) {
		tb := NewTokenBucket(60, 60)
		tb.Tokens = 0
		lastTime, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 00:00:00")
		tb.Last = lastTime
		takeTime, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 00:01:01")

		r := tb.Take(takeTime)

		if r != 59 {
			t.Errorf("got %v, want %v", r, 59)
		}
	})
}
