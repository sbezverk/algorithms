package stack

import (
	"math/rand"
	"testing"
	"time"
)

func TestStackString(t *testing.T) {
	tests := []struct {
		name  string
		input []string
	}{
		{
			name:  "stack of strings",
			input: []string{"string 1", "string 2", "string 3"},
		},
		{
			name:  "stack of 1 string",
			input: []string{"string 1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStackLL[string]()
			if !stack.IsEmpty() {
				t.Fatalf("stack supposed to be empty")
			}
			for _, i := range tt.input {
				stack.Push(i)
			}
			if stack.IsEmpty() {
				t.Fatalf("stack supposed to be non empty")
			}
			empty := false
			l := 0
			for !empty {
				if stack.IsEmpty() {
					empty = true
					continue
				}
				i := stack.Pop()
				l++
				if i != tt.input[len(tt.input)-l] {
					t.Fatalf("expected: %s received: %s", i, tt.input[len(tt.input)-l])
				}
			}
			if l != len(tt.input) {
				t.Fatalf("number of pushed items %d does not match number of popped %d", len(tt.input), l)
			}
		})
	}
}

func TestStackInt(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "stack of ints",
			input: []int{1, 2, 3},
		},
		{
			name:  "stack of 1 int",
			input: []int{100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStackLL[int]()
			if !stack.IsEmpty() {
				t.Fatalf("stack supposed to be empty")
			}
			for _, i := range tt.input {
				stack.Push(i)
			}
			if stack.IsEmpty() {
				t.Fatalf("stack supposed to be non empty")
			}
			empty := false
			l := 0
			for !empty {
				if stack.IsEmpty() {
					empty = true
					continue
				}
				i := stack.Pop()
				l++
				if i != tt.input[len(tt.input)-l] {
					t.Fatalf("expected: %d received: %d", i, tt.input[len(tt.input)-l])
				}
			}
			if l != len(tt.input) {
				t.Fatalf("number of pushed items %d does not match number of popped %d", len(tt.input), l)
			}
		})
	}
}

func TestStackRAString(t *testing.T) {
	tests := []struct {
		name  string
		input []string
	}{
		{
			name:  "stack of strings",
			input: []string{"string 1", "string 2", "string 3"},
		},
		{
			name:  "stack of 1 string",
			input: []string{"string 1"},
		},
		{
			name:  "stack of 2 strings",
			input: []string{"string 1", "string 2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStackRA[string]()
			if !stack.IsEmpty() {
				t.Fatalf("stack supposed to be empty")
			}
			for _, i := range tt.input {
				stack.Push(i)
			}
			if stack.IsEmpty() {
				t.Fatalf("stack supposed to be non empty")
			}
			empty := false
			l := 0
			for !empty {
				if stack.IsEmpty() {
					empty = true
					continue
				}
				i := stack.Pop()
				l++
				if i != tt.input[len(tt.input)-l] {
					t.Fatalf("expected: %s received: %s", i, tt.input[len(tt.input)-l])
				}
			}
			if l != len(tt.input) {
				t.Fatalf("number of pushed items %d does not match number of popped %d", len(tt.input), l)
			}
		})
	}
}

func TestStackRAInt(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "stack of ints",
			input: []int{1, 2, 3},
		},
		{
			name:  "stack of 1 int",
			input: []int{100},
		},
		{
			name:  "stack of 2 ints",
			input: []int{100, 200},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStackRA[int]()
			if !stack.IsEmpty() {
				t.Fatalf("stack supposed to be empty")
			}
			for _, i := range tt.input {
				stack.Push(i)
			}
			if stack.IsEmpty() {
				t.Fatalf("stack supposed to be non empty")
			}
			empty := false
			l := 0
			for !empty {
				if stack.IsEmpty() {
					empty = true
					continue
				}
				i := stack.Pop()
				l++
				if i != tt.input[len(tt.input)-l] {
					t.Fatalf("expected: %d received: %d", i, tt.input[len(tt.input)-l])
				}
			}
			if l != len(tt.input) {
				t.Fatalf("number of pushed items %d does not match number of popped %d", len(tt.input), l)
			}
		})
	}
}

func TestPopEmptyStack(t *testing.T) {
	s := NewStackLL[int]()
	s.Push(1)
	s.Push(3)

	t.Logf("Item 1: %d", s.Pop())
	t.Logf("Item 2: %d", s.Pop())
	t.Logf("Item 3: %d", s.Pop())
	t.Logf("Item 4: %d", s.Pop())
	s = NewStackRA[int]()
	s.Push(1)
	s.Push(3)

	t.Logf("Item 1: %d", s.Pop())
	t.Logf("Item 2: %d", s.Pop())
	t.Logf("Item 3: %d", s.Pop())
	t.Logf("Item 4: %d", s.Pop())
}

func genArrayIntRandom(size int) []int {
	a := make([]int, size)
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		a[i] = -100 + src.Intn(200)
	}

	return a
}

func TestStackPerformance(t *testing.T) {
	ll := NewStackLL[int]()
	ra := NewStackRA[int]()
	size := 100000
	a := genArrayIntRandom(size)

	llPt := time.Now()
	for i := 0; i < size; i++ {
		ll.Push(a[i])
	}
	t.Logf("Link list push time for %d elements was %d microseconds", size, time.Since(llPt).Microseconds())

	raPt := time.Now()
	for i := 0; i < size; i++ {
		ra.Push(a[i])
	}
	t.Logf("Resizing Array push time for %d elements was %d microseconds", size, time.Since(raPt).Microseconds())

	llPot := time.Now()
	for i := 0; i < size; i++ {
		ll.Pop()
	}
	t.Logf("Link list pop time for %d elements was %d microseconds", size, time.Since(llPot).Microseconds())

	raPot := time.Now()
	for i := 0; i < size; i++ {
		ra.Pop()
	}
	t.Logf("Resizing Array pop time for %d elements was %d microseconds", size, time.Since(raPot).Microseconds())
}
