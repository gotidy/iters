package iters

import (
	"fmt"
	"iter"
	"math"
	"slices"
	"testing"
	"time"
)

func printSeq[V any](seq iter.Seq[V]) {
	for v := range seq {
		fmt.Println(v)
	}
}

func printSeq2[K, V any](seq iter.Seq2[K, V]) {
	for k, v := range seq {
		fmt.Println(k, v)
	}
}

func ExampleFilter() {
	printSeq(Filter(
		slices.Values([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
		func(i int) bool { return i%2 == 0 },
	))

	// Output:
	// 0
	// 2
	// 4
	// 6
	// 8
}

func ExampleFilter2() {
	printSeq2(Filter2(
		slices.All([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
		func(_, v int) bool { return v%2 == 0 },
	))

	// Output:
	// 0 0
	// 2 2
	// 4 4
	// 6 6
	// 8 8
}

func ExampleMap() {
	printSeq(Map(
		slices.Values([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
		func(i int) int { return i * 10 },
	))

	// Output:
	// 0
	// 10
	// 20
	// 30
	// 40
	// 50
	// 60
	// 70
	// 80
	// 90
}

func ExampleMap2() {
	printSeq2(Map2(
		slices.All([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
		func(i, v int) (int, int) { return i * 10, v * 100 },
	))

	// Output:
	// 0 0
	// 10 100
	// 20 200
	// 30 300
	// 40 400
	// 50 500
	// 60 600
	// 70 700
	// 80 800
	// 90 900
}

func ExampleMapValues() {
	printSeq2(MapValues(
		slices.All([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
		func(v int) int { return v * 10 },
	))

	// Output:
	// 0 0
	// 1 10
	// 2 20
	// 3 30
	// 4 40
	// 5 50
	// 6 60
	// 7 70
	// 8 80
	// 9 90
}

func ExampleMapKeys() {
	printSeq2(MapKeys(
		slices.All([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
		func(i int) int { return i * 10 },
	))

	// Output:
	// 0 0
	// 10 1
	// 20 2
	// 30 3
	// 40 4
	// 50 5
	// 60 6
	// 70 7
	// 80 8
	// 90 9
}

func ExampleNotNil() {
	printSeq(
		Map(
			NotNil(
				Map(
					slices.Values([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
					func(i int) *int {
						if i%2 == 0 {
							return &i
						}
						return nil
					},
				),
			),
			func(i *int) int { return *i },
		),
	)

	// Output:
	// 0
	// 2
	// 4
	// 6
	// 8
}

func ExampleNotNilValues() {
	printSeq2(
		MapValues(
			NotNilValues(
				MapValues(
					slices.All([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
					func(i int) *int {
						if i%2 == 0 {
							return &i
						}
						return nil
					},
				),
			),
			func(i *int) int { return *i },
		),
	)

	// Output:
	// 0 0
	// 2 2
	// 4 4
	// 6 6
	// 8 8
}

func ExampleNotEmpty() {
	printSeq(
		NotEmpty(
			Map(
				slices.Values([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
				func(i int) int {
					if i%2 == 0 {
						return i
					}
					return 0
				},
			),
		),
	)

	// Output:
	// 2
	// 4
	// 6
	// 8
}

func ExampleNotEmptyValues() {
	printSeq2(
		NotEmptyValues(
			MapValues(
				slices.All([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
				func(i int) int {
					if i%2 == 0 {
						return i
					}
					return 0
				},
			),
		),
	)

	// Output:
	// 2 2
	// 4 4
	// 6 6
	// 8 8
}

func ExampleWithKeys() {
	printSeq2(WithKeys(
		slices.Values([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
		func(v int) int { return v * 10 },
	))

	// Output:
	// 0 0
	// 10 1
	// 20 2
	// 30 3
	// 40 4
	// 50 5
	// 60 6
	// 70 7
	// 80 8
	// 90 9
}

func ExampleToSeq2() {
	printSeq2(ToSeq2(
		slices.Values([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
		func(v int) (int, int) { return v * 10, v * 100 },
	))

	// Output:
	// 0 0
	// 10 100
	// 20 200
	// 30 300
	// 40 400
	// 50 500
	// 60 600
	// 70 700
	// 80 800
	// 90 900
}

func ExampleFold() {
	printSeq(Fold(
		slices.Values([]int{10, 0, 1, 1, 2, 2, 5, 5, 7, 9, 9, 10}),
	))

	// Output:
	// 10
	// 0
	// 1
	// 2
	// 5
	// 7
	// 9
}

func ExampleFold2() {
	printSeq2(Fold2(
		Map2(
			slices.All([]int{10, 0, 1, 1, 2, 2, 5, 5, 7, 9, 9, 10}),
			func(k, v int) (int, int) { return v, k }, // Swap kes and values.
		),
	))

	// Output:
	// 10 0
	// 0 1
	// 1 2
	// 2 4
	// 5 6
	// 7 8
	// 9 9
}

func ExampleFold2Func() {
	printSeq2(Fold2Func(
		slices.All([]int{10, 0, 1, 1, 2, 2, 5, 5, 7, 9, 9, 10}),
		func(_, v int) int { return v },
	),
	)

	// Output:
	// 0 10
	// 1 0
	// 2 1
	// 4 2
	// 6 5
	// 8 7
	// 9 9
}

func ExampleReduce() {
	fmt.Println(Reduce(
		slices.Values([]int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1}),
		100,
		func(r int, v int) int { return r + v },
	))

	// Output:
	// 105
}

func ExampleValues() {
	printSeq(Values(
		slices.All([]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}),
	))

	// Output:
	// 0
	// 10
	// 20
	// 30
	// 40
	// 50
	// 60
	// 70
	// 80
	// 90
}

func ExampleKeys() {
	printSeq(Keys(
		slices.All([]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}),
	))

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
}

func ExampleContains() {
	s := slices.Values([]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90})
	fmt.Println(Contains(50, s))
	fmt.Println(Contains(31, s))
	// Output:
	// true
	// false
}

func ExampleEqual() {
	s1 := slices.Values([]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90})
	s2 := slices.Values([]int{0, 10, 20, 30, 140, 50, 60, 70, 80, 90})
	s3 := slices.Values([]int{0, 10, 20})
	s4 := slices.Values([]int{})
	s5 := slices.Values([]int(nil))
	fmt.Println(Equal(s1, s1))
	fmt.Println(Equal(s1, s2))
	fmt.Println(Equal(s1, s3))
	fmt.Println(Equal(s3, s3))
	fmt.Println(Equal(s4, s4))
	fmt.Println(Equal(s5, s5))
	fmt.Println(Equal(s5, s1))
	// Output:
	// true
	// false
	// false
	// true
	// true
	// true
	// false
}

func ExampleMerge() {
	s1 := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	s2 := []int{90, 80, 70, 60, 50, 40, 30, 20, 10, 0}

	seq1 := slices.Values(s1)
	seq2 := slices.Values(s2)

	fmt.Println(Equal(Merge(seq1), seq1))
	fmt.Println(Equal(Merge(seq1, seq2), slices.Values(append(s1, s2...))))

	printSeq(Merge[int]())

	// Output:
	// true
	// true
}

func ExampleMerge2() {
	s1 := slices.All([]int{0, 10, 20, 30, 40})
	s2 := slices.All([]int{50, 60, 70, 80, 90})

	printSeq2(Merge2[int, int]())
	printSeq2(Merge2(s1))
	printSeq2(Merge2(s1, s2))
	// Output:
	// 0 0
	// 1 10
	// 2 20
	// 3 30
	// 4 40
	// 0 0
	// 1 10
	// 2 20
	// 3 30
	// 4 40
	// 0 50
	// 1 60
	// 2 70
	// 3 80
	// 4 90
}

func BenchmarkEqual(b *testing.B) {
	s := make([]int, 100)
	for _, i := range s {
		s[i] = i
	}
	seq := slices.Values(s)
	b.ResetTimer()
	b.Run("iterator equality", func(b *testing.B) {
		for range b.N {
			Equal(seq, seq)
		}
	})
	b.Run("slice equality", func(b *testing.B) {
		for range b.N {
			slices.Equal(s, s)
		}
	})
}

func ExampleCount() {
	s := slices.Values([]int{1, 2, 3, 4, 5})

	fmt.Println(Count(s))
	fmt.Println(CountFunc(s, func(v int) bool { return v%2 == 0 }))
	// Output:
	// 5
	// 2
}

func ExampleCount2() {
	s := slices.All([]int{2, 4, 5, 8, 10})

	fmt.Println(Count2(s))
	fmt.Println(CountFunc2(s, func(k, v int) bool { return k%2 == 0 && v%2 == 0 }))
	// Output:
	// 5
	// 2
}

func ExampleGroup() {
	var testSeq iter.Seq2[string, int] = func(yield func(string, int) bool) {
		yield("a", 1)
		yield("b", 2)
		yield("a", 3)
		yield("c", 4)
		yield("b", 5)
	}

	result := Group(testSeq)

	fmt.Println(result["a"])
	fmt.Println(result["b"])
	fmt.Println(result["c"])

	// Output:
	// [1 3]
	// [2 5]
	// [4]
}

func ExampleGroupFunc() {
	var testSeq iter.Seq[int] = func(yield func(int) bool) {
		yield(1)
		yield(2)
		yield(3)
		yield(4)
		yield(5)
	}

	result := GroupFunc(testSeq, func(v int) string {
		if v%2 == 0 {
			return "even"
		}
		return "odd"
	})

	fmt.Println("Even:", result["even"])
	fmt.Println("Odd:", result["odd"])

	// Output:
	// Even: [2 4]
	// Odd: [1 3 5]
}

func ExamplePointers() {
	for v := range Pointers([]int{5, 7, 1}) {
		fmt.Println(*v)
	}

	for v := range Pointers([]int(nil)) {
		fmt.Println(*v)
	}

	// Output:
	// 5
	// 7
	// 1
}

func ExampleExponential() {
	start := 1.0
	max := 100.0
	factor := 2.0

	fmt.Println(slices.Collect(Trim(Exponential(start, max, factor), 10)))

	// Output:
	// [1 2 4 8 16 32 64 100 100 100]
}

func ExampleRepeat() {
	fmt.Println(slices.Collect(Trim(Repeat(11), 10)))

	// Output:
	// [11 11 11 11 11 11 11 11 11 11]
}

func ExampleOf() {
	fmt.Println(slices.Collect(Of(1, 2, 4, 8, 16, 32, 64)))

	// Output:
	// [1 2 4 8 16 32 64]
}

func Test_jitter(t *testing.T) {
	t.Parallel()

	assertEquals(t, 1, jitter(2, 0.5, -1.0))
	assertEquals(t, 2, jitter(2, 0.5, 0.0))
	assertEquals(t, 3, jitter(2, 0.5, 1.0))
}

func TestJitter(t *testing.T) {
	t.Parallel()

	seq := Trim(Repeat(20.0), 1000)

	t.Run("factor 0", func(t *testing.T) {
		for v := range Jitter(seq, 0.0) {
			assertEquals(t, v, 20.0)
		}
	})

	t.Run("factor 0.1", func(t *testing.T) {
		jitteredSeq := Jitter(seq, 0.1)
		if !isUniform(slices.Collect(jitteredSeq), 18.0, 22) {
			t.Error("Jitter is not uniform")
		}
	})
}

func assertEquals[T comparable](t *testing.T, expected, actual T) {
	if expected != actual {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

// uniformity calculates the uniformity of values in a given range.
func isUniform(values []float64, minValue, maxValue float64) bool {
	const bins = 10
	histogram := make([]int, bins)
	binWidth := (maxValue - minValue) / float64(bins)
	for _, value := range values {
		if value >= minValue && value < maxValue {
			histogram[min(int((value-minValue)/binWidth), bins-1)]++
		}
	}

	expected := float64(len(values)) / float64(bins)
	var square float64
	for _, count := range histogram {
		square += math.Pow(float64(count)-expected, 2) / expected
	}

	return square < 16.919
}

func ExampleMaxElapsedTime() {
	start := time.Now()
	for range MaxElapsedTime(Repeat(10), time.Millisecond*100) {
		time.Sleep(time.Millisecond)
	}
	fmt.Println(time.Since(start) > time.Millisecond*100)

	// Output:
	// true
}
