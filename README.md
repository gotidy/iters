[![CI State](https://github.com/gotidy/iters/actions/workflows/Go/badge.svg)](https://github.com/gotidy/iters/actions)
[![Go Doc](https://godoc.org/github.com/gotidy/iters?status.svg)](https://pkg.go.dev/github.com/gotidy/iters)

# iters

Go iterators.

## Installation

```go
go get github.com/gotidy/iters
```

## General Iterators

The `iters` library provides a set of general-purpose iterators that simplify the process of transformation, creation and using collections.

Examples:

**Filter**: Filtering sequence elements.

```go
seq := Filter(
    slices.Values([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
    func(i int) bool { return i%2 == 0 },
)   
```

**Map**: Transforming sequence elements

```go
for v := range Map(slices.Values([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}), func(i int) string { return strconv.Itoa(i) }) {
    ...
}
```

Other examples can be found in tests.

## Retry

The `Retry` iterator is allowed to iterate over sequence of delays, with the specified delays. The `Retry` waits for the specified delay before retrying, except cases when context is cancelled.

The `Repeat`, `Trim`, `Of`, `Exponential`, `Jitter`, `MaxElapsedTime` functions can be used to define delays suppliers. It is also possible to define your own iterator for special behavior.

```go
for attempt, delay := range Retry(context.Background(), Jitter(Trim(Exponential(time.Millisecond, time.Second, 2), 5), 0.5)) {
    if err = doSomething(); err == nil {
        return nil
    }
    fmt.Println(attempt, delay)
}
return err
```

```go
for range MaxElapsedTime(Retry(ctx, Repeat(time.Second)), time.Minute) {
    ...
}
```
