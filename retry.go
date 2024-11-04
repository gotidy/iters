package iters

import (
	"context"
	"iter"
	"time"
)

// Retry returns sequence that allows to retry with specified delays.
func Retry[R int, D time.Duration](ctx context.Context, delays iter.Seq[time.Duration]) iter.Seq2[int, time.Duration] {
	return func(yield func(int, time.Duration) bool) {
		attempts := 1
		for delay := range delays {
			if !yield(attempts, delay) {
				return
			}
			select {
			case <-ctx.Done():
				return
			case <-time.After(delay):
			}
			attempts++
		}
	}
}
