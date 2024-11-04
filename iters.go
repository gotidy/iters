package iters

import (
	"iter"
	"math"
	"math/rand/v2"
	"time"

	"golang.org/x/exp/constraints"
)

// Filter filters values from the the sequence ov values using a filter function.
func Filter[V any](seq iter.Seq[V], f func(item V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if f(v) && !yield(v) {
				return
			}
		}
	}
}

// Filter2 filters values from the the sequence ov key-values pairs using a filter function.
func Filter2[K, V any](seq iter.Seq2[K, V], f func(k K, v V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if f(k, v) && !yield(k, v) {
				return
			}
		}
	}
}

// Map converts the sequence of values to the sequence of values another type using a mapping function.
func Map[V1, V2 any](seq iter.Seq[V1], f func(v V1) V2) iter.Seq[V2] {
	return func(yield func(V2) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

// Map2 converts the sequence of values to the sequence of values another type using a mapping function.
func Map2[K1, K2, V1, V2 any](seq iter.Seq2[K1, V1], f func(k K1, v V1) (K2, V2)) iter.Seq2[K2, V2] {
	return func(yield func(K2, V2) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

// MapValues converts the sequence of key-value pairs using a values mapping function.
func MapValues[K, V1, V2 any](seq iter.Seq2[K, V1], f func(item V1) V2) iter.Seq2[K, V2] {
	return func(yield func(K, V2) bool) {
		for k, v := range seq {
			if !yield(k, f(v)) {
				return
			}
		}
	}
}

// MapKeys converts the sequence of key-value pairs using a keys mapping function.
func MapKeys[K1, K2, V any](seq iter.Seq2[K1, V], f func(item K1) K2) iter.Seq2[K2, V] {
	return func(yield func(K2, V) bool) {
		for k, v := range seq {
			if !yield(f(k), v) {
				return
			}
		}
	}
}

// NotNil skips nil values in the sequence.
func NotNil[V any, P *V](seq iter.Seq[P]) iter.Seq[P] {
	return Filter(seq, func(p P) bool {
		return p != nil
	})
}

// NotNilValues skips nil values in the sequence.
func NotNilValues[K, V any, P *V](seq iter.Seq2[K, P]) iter.Seq2[K, P] {
	return Filter2(seq, func(_ K, p P) bool {
		return p != nil
	})
}

// NotEmpty skips zero values in the sequence.
func NotEmpty[V comparable](seq iter.Seq[V]) iter.Seq[V] {
	return Filter(seq, func(v V) bool {
		var zero V
		return v != zero
	})
}

// NotEmptyValues skips zero values in the sequence.
func NotEmptyValues[K any, V comparable](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return Filter2(seq, func(_ K, v V) bool {
		var zero V
		return v != zero
	})
}

// WithKeys converts the sequence of values to the sequence of key-value pairs by adding key to the sequence.
func WithKeys[K, V any](seq iter.Seq[V], f func(item V) K) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for v := range seq {
			if !yield(f(v), v) {
				return
			}
		}
	}
}

// WithIndex converts a sequence of values into a key-value sequence, where key is an index starting with 0.
func WithIndex[V any](seq iter.Seq[V]) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		i := 0
		for v := range seq {
			if !yield(i, v) {
				return
			}
			i++
		}
	}
}

// ToSeq2 converts the sequence of of individual values to the sequence of key-value pairs.
func ToSeq2[T, K, V any](seq iter.Seq[T], f func(item T) (K, V)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

// Fold push items from one sequence to another sequence skipping duplicates.
func Fold[V comparable](seq iter.Seq[V]) iter.Seq[V] {
	return FoldFunc(seq, func(v V) V { return v })
}

// FoldFunc push items from one sequence to another sequence skipping duplicates.
func FoldFunc[K comparable, V any](seq iter.Seq[V], foldKey func(V) K) iter.Seq[V] {
	return func(yield func(V) bool) {
		m := make(map[K]struct{})
		for v := range seq {
			key := foldKey(v)
			if _, ok := m[key]; ok {
				continue
			}
			m[key] = struct{}{}
			if !yield(v) {
				return
			}
		}
	}
}

// Fold2 push items from one sequence to another sequence skipping duplicates.
func Fold2[K comparable, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		m := make(map[K]struct{})
		for k, v := range seq {
			if _, ok := m[k]; ok {
				continue
			}
			m[k] = struct{}{}
			if !yield(k, v) {
				return
			}
		}
	}
}

// Fold2Func push items from one sequence to another sequence skipping duplicates.
func Fold2Func[F comparable, K, V any](seq iter.Seq2[K, V], foldKey func(K, V) F) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		m := make(map[F]struct{})
		for k, v := range seq {
			key := foldKey(k, v)
			if _, ok := m[key]; ok {
				continue
			}
			m[key] = struct{}{}
			if !yield(k, v) {
				return
			}
		}
	}
}

// Reduce reduces a sequence to a single value using a reduction function.
func Reduce[T, R any](seq iter.Seq[T], initializer R, f func(R, T) R) R {
	r := initializer
	for v := range seq {
		r = f(r, v)
	}

	return r
}

// Values convert Seq2 to a Seq by returning the values of the sequence.
func Values[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range seq {
			if !yield(v) {
				return
			}
		}
	}
}

// Keys convert Seq2 to a Seq by returning the keys of the sequence.
func Keys[K, V any](seq iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range seq {
			if !yield(k) {
				return
			}
		}
	}
}

// Contains checks that the sequence contains the specified value.
func Contains[T comparable](s T, in iter.Seq[T]) bool {
	for v := range in {
		if s == v {
			return true
		}
	}
	return false
}

// Equal compare two sequences. Slow.
func Equal[T comparable](s1, s2 iter.Seq[T]) bool {
	return EqualFunc(s1, s2, func(v1, v2 T) bool { return v1 == v2 })
}

// EqualFunc compare two sequences. Slow.
func EqualFunc[T1, T2 any](s1 iter.Seq[T1], s2 iter.Seq[T2], equal func(T1, T2) bool) bool {
	next1, stop1 := iter.Pull(s1)
	next2, stop2 := iter.Pull(s2)
	for {
		v1, ok1 := next1()
		v2, ok2 := next2()
		if ok1 != ok2 || (ok1 && ok2 && !equal(v1, v2)) {
			stop1()
			stop2()
			return false
		}
		if !ok1 {
			stop1()
			stop2()
			return true
		}
	}
}

// Merge sequences of values into one.
func Merge[V any](seqs ...iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, seq := range seqs {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Merge2 sequences of key-value pairs into one.
func Merge2[K, V any](seqs ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq := range seqs {
			for k, v := range seq {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// Count values.
func Count[V any](s iter.Seq[V]) int {
	var count int
	for range s {
		count++
	}
	return count
}

// CountFunc values.
func CountFunc[V any](s iter.Seq[V], f func(V) bool) int {
	var count int
	for v := range s {
		if f(v) {
			count++
		}
	}
	return count
}

// Count values.
func Count2[K, V any](s iter.Seq2[K, V]) int {
	var count int
	for range s {
		count++
	}
	return count
}

// CountFunc2 counts values.
func CountFunc2[K, V any](s iter.Seq2[K, V], f func(K, V) bool) int {
	var count int
	for k, v := range s {
		if f(k, v) {
			count++
		}
	}
	return count
}

// Group group sequence by key.
func Group[K comparable, V any](seq iter.Seq2[K, V]) map[K][]V {
	m := make(map[K][]V)
	for k, v := range seq {
		m[k] = append(m[k], v)
	}
	return m
}

// GroupFunc group sequence by key.
func GroupFunc[K comparable, V any](seq iter.Seq[V], key func(V) K) map[K][]V {
	return Group(WithKeys(seq, key))
}

// Pointers returns a sequence of pointers to the elements of the given slice.
func Pointers[V any](vv []V) iter.Seq[*V] {
	return func(yield func(*V) bool) {
		for i := range vv {
			if !yield(&vv[i]) {
				return
			}
		}
	}
}

// Repeat repeats a value infinitely.
func Repeat[V any](v V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for {
			if !yield(v) {
				return
			}
		}
	}
}

// Trim trims a sequence by count.
func Trim[V any](s iter.Seq[V], count int) iter.Seq[V] {
	return func(yield func(V) bool) {
		n := 0
		for v := range s {
			if n >= count {
				return
			}
			if !yield(v) {
				return
			}
			n++
		}
	}
}

// Of creates a sequence of members.
func Of[V any](vv ...V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range vv {
			if !yield(v) {
				return
			}
		}
	}
}

// Exponential generate exponential sequence of values.
// The first value will be start, each following v = v * factor, but not greater than max.
func Exponential[D constraints.Float | constraints.Integer](start, max D, factor float64) iter.Seq[D] {
	return func(yield func(D) bool) {
		v := start
		for {
			if !yield(v) {
				return
			}
			if v < max {
				v = min(max, D(float64(v)*factor))
			}
		}
	}
}

func jitter[T constraints.Float | constraints.Integer](v T, factor, random float64) T {
	if factor == 0 {
		return v
	}
	delta := factor * float64(v)

	// Get a random value from the range [minInterval, maxInterval].
	// The formula used below has a +1 because if the minInterval is 1 and the maxInterval is 3 then
	// we want a 33% chance for selecting either 1, 2 or 3.
	return T(float64(v) + delta*random)
}

// Jitter returns sequence with added jitter to the values.
// value = value * (random value in range [1 - Jitter, 1 + Jitter]).
// Example: jitter of 10 with factor 0.1 will returns values in range [9, 11].
func Jitter[T constraints.Float | constraints.Integer](vv iter.Seq[T], factor float64) iter.Seq[T] {
	return func(yield func(T) bool) {
		jitter := func(v T) T {
			return jitter(v, factor, float64(rand.Uint64())*2/math.MaxUint64-1)
		}
		if factor == 0 {
			jitter = func(v T) T { return v }
		}
		for v := range vv {
			if !yield(jitter(v)) {
				return
			}
		}
	}
}

// MaxElapsedTime stops sequence processing after the specified time has elapsed.
func MaxElapsedTime[T any](seq iter.Seq[T], max time.Duration) iter.Seq[T] {
	return func(yield func(T) bool) {
		start := time.Now()
		for v := range seq {
			if time.Since(start) > max {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}
