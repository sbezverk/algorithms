package stack

import (
	"testing"
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
			stack := NewStack[string]()
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
			stack := NewStack[int]()
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
