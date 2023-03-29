package quick_union_weighted_pathcompression

import "fmt"

type QUWPC interface {
	Union(int, int) error
	Connected(int, int) (bool, error)
	Root(int) (int, error)
	Opened(int) (bool, error)
	String() string
}

type site struct {
	// By default a site is closed
	opened bool
}

type quwpc struct {
	sites []*site
	ids   []int
	sz    []int
}

func (qu *quwpc) String() string {
	s := ""
	for i := 0; i < len(qu.ids); i++ {

	}

	return s
}

func (qu *quwpc) Union(p, q int) error {
	if p < 0 || q < 0 {
		return fmt.Errorf("site ids cannot be negative number, accepted values are from 1 to %d", len(qu.sites))
	}
	if p+1 > len(qu.sites) || q+1 > len(qu.sites) {
		return fmt.Errorf("site ids exceeding the boundaries, accepted values are from 1 to %d", len(qu.sites))
	}
	i, _ := qu.Root(p)
	j, _ := qu.Root(q)
	switch {
	case i == j:
	case qu.sz[i] < qu.sz[j]:
		qu.ids[i] = j
		qu.sz[j] += qu.sz[i]
	default:
		// qu.sz[i] > qu.sz[j]
		qu.ids[j] = i
		qu.sz[i] += qu.sz[j]
	}

	return nil
}

func (qu *quwpc) Connected(p, q int) (bool, error) {
	if p < 0 || q < 0 {
		return false, fmt.Errorf("site ids cannot be negative number, accepted values are from 1 to %d", len(qu.sites))
	}
	if p+1 > len(qu.sites) || q+1 > len(qu.sites) {
		return false, fmt.Errorf("site ids exceeding the boundaries, accepted values are from 1 to %d", len(qu.sites))
	}
	rp, _ := qu.Root(p)
	rq, _ := qu.Root(q)
	return rq == rp, nil
}

func (qu *quwpc) Root(p int) (int, error) {
	if p+1 > len(qu.ids) {
		return 0, fmt.Errorf("site id exceeding the boundaries, accepted values are from 1 to %d", len(qu.ids))
	}
	found := false
	for !found && p < len(qu.ids) {
		if p == qu.ids[p] {
			return p, nil
		}
		// Path compression
		qu.ids[p] = qu.ids[qu.ids[p]]
		p = qu.ids[p]
	}

	return 0, fmt.Errorf("root for site %d is not found, it is a bug", p)
}

func (qu *quwpc) Opened(p int) (bool, error) {
	if p > len(qu.sites) {
		return false, fmt.Errorf("site %d is outside of the boundaries", p)
	}
	return qu.sites[p].opened, nil
}

func NewQUWPC(n int) (QUWPC, error) {
	if n < 2 {
		return nil, fmt.Errorf("invalid value of the number of sites: %d", n)
	}
	s := make([]*site, n)
	ids := make([]int, n)
	sz := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = &site{}
		ids[i] = i
		sz[i] = 1
	}

	return &quwpc{
		sites: s,
		ids:   ids,
		sz:    sz,
	}, nil
}
