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

func TestTake(t *testing.T) {
	t.Run("When token is zero must return zero and not minus one", func(t *testing.T) {
		tb := NewTokenBucket(0, 60)
		takeTime, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 00:00:00")
		r := tb.Take(takeTime)
		if r != 0 {
			t.Errorf("got %v, want 0", r)
		}
		if tb.Tokens < 0 {
			t.Errorf("got %v, want 0", tb.Tokens)
		}
	})

	t.Run("Taking one token and check tokens will decrease one token", func(t *testing.T) {
		tb := NewTokenBucket(1, 60)
		takeTime, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 00:00:00")
		r := tb.Take(takeTime)
		if r != 1 {
			t.Errorf("got %v, want 0", r)
		}
		if tb.Tokens != 0 {
			t.Errorf("got %v, want 0", tb.Tokens)
		}
		if tb.Last != takeTime {
			t.Errorf("got %v, want 0", tb.Last)
		}
	})
}
