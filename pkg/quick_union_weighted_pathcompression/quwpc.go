package quick_union_weighted_pathcompression

import "fmt"

type QUWPC interface {
	Union(int, int)
	Connected(int, int) bool
	Root(int) int
	Opened(int) (bool, error)
}

type site struct {
	// By default a site is closed
	opened bool
}

type quwpc struct {
	sites []*site
	ids   []int
}

func (qu *quwpc) Union(p, q int) {

}

func (qu *quwpc) Connected(p, q int) bool {

	return false
}

func (qu *quwpc) Root(p int) int {

	return 0
}

func (qu *quwpc) Opened(p int) (bool, error) {
	if p > len(qu.sites) {
		return false, fmt.Errorf("site %d is outside of the boundaries", p)
	}
	return qu.sites[p].opened, nil
}

func NewQUWPC(n int) (QUWPC, error) {
	if n < 0 {
		return nil, fmt.Errorf("invalid value of the number of sites: %d", n)
	}

	return &quwpc{
		sites: make([]*site, n),
		ids:   make([]int, n),
	}, nil
}
