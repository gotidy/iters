package iters

import (
	"context"
	"iter"
	"time"
)

// Retry returns sequence that allows to retry with specified delays.
// The first iteration occurs immediately (retry, delay, retry, delay, retry).
func Retry[R int, D time.Duration](ctx context.Context, delays iter.Seq[time.Duration]) iter.Seq2[int, time.Duration] {
	return func(yield func(int, time.Duration) bool) {
		if !yield(0, 0) {
			return
		}
		for attempt, delay := range RetryAfterDelay(ctx, delays) {
			if !yield(attempt, delay) {
				return
			}
		}
	}
}

// RetryAfterDelay returns sequence that allows to retry with specified delays.
// Started from delay (delay, retry, delay, retry).
func RetryAfterDelay[R int, D time.Duration](ctx context.Context, delays iter.Seq[time.Duration]) iter.Seq2[int, time.Duration] {
	return func(yield func(int, time.Duration) bool) {
		if delays == nil {
			return
		}
		attempts := 1
		for delay := range delays {
			select {
			case <-ctx.Done():
				return
			case <-time.After(delay):
			}
			if !yield(attempts, delay) {
				return
			}
			attempts++
		}
	}
}
