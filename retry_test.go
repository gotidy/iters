package iters

import (
	"context"
	"fmt"
	"time"
)

func ExampleRetry() {
	start := time.Now()
	for attempt, delay := range Retry(context.Background(), Trim(Repeat(time.Millisecond*100), 10)) {
		fmt.Println(attempt, delay)
	}
	fmt.Println(time.Since(start) > time.Second)

	// Output:
	// 0 0s
	// 1 100ms
	// 2 100ms
	// 3 100ms
	// 4 100ms
	// 5 100ms
	// 6 100ms
	// 7 100ms
	// 8 100ms
	// 9 100ms
	// 10 100ms
	// true
}

func ExampleRetry_break() {
	for attempt, delay := range Retry(context.Background(), Trim(Exponential(time.Millisecond, time.Second, 2), 5)) {
		fmt.Println(attempt, delay)
		if attempt == 5 {
			break
		}
	}

	// Output:
	// 0 0s
	// 1 1ms
	// 2 2ms
	// 3 4ms
	// 4 8ms
	// 5 16ms
}

func ExampleRetry_ctx() {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		for range Retry(ctx, Repeat(time.Millisecond*100)) {
		}
		fmt.Println("stopped")
		close(done)
	}()

	cancel()
	<-done

	// Output:
	// stopped
}
