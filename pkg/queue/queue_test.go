package queue

import (
	"math/rand"
	"testing"
	"time"
)

func TestQueueString(t *testing.T) {
	tests := []struct {
		name  string
		input []string
	}{
		{
			name:  "queue of strings",
			input: []string{"string 1", "string 2", "string 3"},
		},
		{
			name:  "queue of 1 string",
			input: []string{"string 1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewQueue[string]()
			if !queue.IsEmpty() {
				t.Fatalf("queue supposed to be empty")
			}
			for _, i := range tt.input {
				queue.Enqueue(i)
			}
			if queue.IsEmpty() {
				t.Fatalf("queue supposed to be non empty")
			}
			empty := false
			l := 0
			for !empty {
				if queue.IsEmpty() {
					empty = true
					continue
				}
				i := queue.Dequeue()
				if i != tt.input[l] {
					t.Fatalf("expected: %s received: %s", i, tt.input[l])
				}
				l++
			}
			if l != len(tt.input) {
				t.Fatalf("number of pushed items %d does not match number of popped %d", len(tt.input), l+1)
			}
		})
	}
}

func TestQueueInt(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "queue of ints",
			input: []int{1, 2, 3},
		},
		{
			name:  "queue of 1 int",
			input: []int{100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewQueue[int]()
			if !queue.IsEmpty() {
				t.Fatalf("queue supposed to be empty")
			}
			for _, i := range tt.input {
				queue.Enqueue(i)
			}
			if queue.IsEmpty() {
				t.Fatalf("queue supposed to be non empty")
			}
			empty := false
			l := 0
			for !empty {
				if queue.IsEmpty() {
					empty = true
					continue
				}
				i := queue.Dequeue()
				if i != tt.input[l] {
					t.Fatalf("expected: %d received: %d", i, tt.input[l])
				}
				l++
			}
			if l != len(tt.input) {
				t.Fatalf("number of pushed items %d does not match number of popped %d", len(tt.input), l)
			}
		})
	}
}

func TestQueueRAString(t *testing.T) {
	tests := []struct {
		name  string
		input []string
	}{
		{
			name:  "queue of strings",
			input: []string{"string 1", "string 2", "string 3"},
		},
		{
			name:  "queue of 1 string",
			input: []string{"string 1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewQueueRA[string]()
			if !queue.IsEmpty() {
				t.Fatalf("queue supposed to be empty")
			}
			for _, i := range tt.input {
				queue.Enqueue(i)
			}
			if queue.IsEmpty() {
				t.Fatalf("queue supposed to be non empty")
			}
			empty := false
			l := 0
			for !empty {
				if queue.IsEmpty() {
					empty = true
					continue
				}
				i := queue.Dequeue()
				if i != tt.input[l] {
					t.Fatalf("expected: %s received: %s", i, tt.input[l])
				}
				l++
			}
			if l != len(tt.input) {
				t.Fatalf("number of pushed items %d does not match number of popped %d", len(tt.input), l+1)
			}
		})
	}
}

func TestQueueRAInt(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "queue of ints",
			input: []int{1, 2, 3},
		},
		{
			name:  "queue of 1 int",
			input: []int{100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewQueueRA[int]()
			if !queue.IsEmpty() {
				t.Fatalf("queue supposed to be empty")
			}
			for _, i := range tt.input {
				queue.Enqueue(i)
			}
			if queue.IsEmpty() {
				t.Fatalf("queue supposed to be non empty")
			}
			empty := false
			l := 0
			for !empty {
				if queue.IsEmpty() {
					empty = true
					continue
				}
				i := queue.Dequeue()
				if i != tt.input[l] {
					t.Fatalf("expected: %d received: %d", i, tt.input[l])
				}
				l++
			}
			if l != len(tt.input) {
				t.Fatalf("number of pushed items %d does not match number of popped %d", len(tt.input), l)
			}
		})
	}
}

func TestDequeueEmptyQueueInt(t *testing.T) {
	s := NewQueue[int]()
	s.Enqueue(1)
	s.Enqueue(3)

	t.Logf("Item 1: %d", s.Dequeue())
	t.Logf("Item 2: %d", s.Dequeue())
	t.Logf("Item 3: %d", s.Dequeue())
	t.Logf("Item 4: %d", s.Dequeue())
}

func TestDequeueEmptyQueueString(t *testing.T) {
	s := NewQueue[string]()
	s.Enqueue("string 1")
	s.Enqueue("string 2")

	t.Logf("Item 1: %s", s.Dequeue())
	t.Logf("Item 2: %s", s.Dequeue())
	t.Logf("Item 3: %s", s.Dequeue())
	t.Logf("Item 4: %s", s.Dequeue())
}

type testObj struct {
	name string
	size int
}

func TestDequeueEmptyQueueObj(t *testing.T) {
	s := NewQueue[*testObj]()
	s.Enqueue(&testObj{name: "string 1", size: 1})
	s.Enqueue(&testObj{name: "string 1", size: 1})

	t.Logf("Item 1: %+v", s.Dequeue())
	t.Logf("Item 2: %+v", s.Dequeue())
	t.Logf("Item 3: %+v", s.Dequeue())
	t.Logf("Item 4: %+v", s.Dequeue())
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
	ll := NewQueue[int]()
	ra := NewQueueRA[int]()
	size := 100000
	a := genArrayIntRandom(size)

	llEt := time.Now()
	for i := 0; i < size; i++ {
		ll.Enqueue(a[i])
	}
	t.Logf("Link list enqueue time for %d elements was %d microseconds", size, time.Since(llEt).Microseconds())

	raEt := time.Now()
	for i := 0; i < size; i++ {
		ra.Enqueue(a[i])
	}
	t.Logf("Resizing Array enqueue time for %d elements was %d microseconds", size, time.Since(raEt).Microseconds())

	llDt := time.Now()
	for i := 0; i < size; i++ {
		ll.Dequeue()
	}
	t.Logf("Link list dequeue time for %d elements was %d microseconds", size, time.Since(llDt).Microseconds())

	raDt := time.Now()
	for i := 0; i < size; i++ {
		ra.Dequeue()
	}
	t.Logf("Resizing Array dequeue time for %d elements was %d microseconds", size, time.Since(raDt).Microseconds())
}
