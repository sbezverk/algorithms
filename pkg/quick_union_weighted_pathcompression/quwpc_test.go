package quick_union_weighted_pathcompression

import (
	"testing"
)

type union struct {
	p int
	q int
}
type input struct {
	n         int
	unions    []*union
	connected []*union
}

func TestQUWPC(t *testing.T) {
	tests := []struct {
		name  string
		input input
		fail  bool
	}{
		{
			name: "2 elements",
			input: input{
				n: 2,
				unions: []*union{{
					p: 1, q: 0},
				},
				connected: []*union{{
					p: 0, q: 1},
				},
			},
		},
		{
			name: "3 elements",
			input: input{
				n: 5,
				unions: []*union{
					{p: 1, q: 0},
					{p: 4, q: 0},
				},
				connected: []*union{
					{p: 0, q: 1},
					{p: 1, q: 4},
				},
			},
		},
		{
			name: "9 elements",
			input: input{
				n: 9,
				unions: []*union{
					{p: 1, q: 0},
					{p: 4, q: 1},
					{p: 7, q: 8},
					{p: 6, q: 7},
					{p: 5, q: 6},
				},
				connected: []*union{
					{p: 0, q: 1},
					{p: 1, q: 4},
					{p: 5, q: 8},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu, err := NewQUWPC(tt.input.n)
			if err != nil && !tt.fail {
				t.Fatalf("supposed to succeed but failed with error: %+v", err)
			}
			if err == nil && tt.fail {
				t.Fatalf("supposed to fail but succeeded")
			}
			if err != nil {
				return
			}
			// Adding required unions
			for _, u := range tt.input.unions {
				err := qu.Union(u.p, u.q)
				if err != nil && !tt.fail {
					t.Fatalf("supposed to succeed but failed with error: %+v", err)
				}
				if err == nil && tt.fail {
					t.Fatalf("supposed to fail but succeeded")
				}
				if err != nil {
					return
				}
			}
			// Check built unions
			for _, u := range tt.input.connected {
				c, err := qu.Connected(u.p, u.q)
				if err != nil && !tt.fail {
					t.Fatalf("supposed to succeed but failed with error: %+v", err)
				}
				if err == nil && tt.fail {
					t.Fatalf("supposed to fail but succeeded")
				}
				if err != nil {
					return
				}
				if !c {
					t.Fatalf("Union %d and %d is supposed to be connected", u.p, u.q)
				}
			}
		})
	}
}
