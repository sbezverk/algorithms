package binary_search

import (
	"math/rand"
	"testing"
	"time"
)

func makeSliceOfRandomInts(l int) []int {
	a := make([]int, l)
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		a[i] = -100 + src.Intn(200)
	}

	return a
}

func RandomInt() int {
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	return -100 + src.Intn(200)
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name  string
		input int
	}{
		{
			name:  "1k ints",
			input: 1024,
		},
		{
			name:  "2k ints",
			input: 2048,
		},
		{
			name:  "4k ints",
			input: 4096,
		},
		{
			name:  "8k ints",
			input: 8192,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := makeSliceOfRandomInts(tt.input)
			s := time.Now()
			e := RandomInt()
			found := BinarySearch(e, a)
			computationTime := time.Since(s)
			t.Logf("Element %d is found %t in array of %d elements, it took %d milliseconds", e, found, len(a), computationTime/time.Millisecond)
		})
	}
}
