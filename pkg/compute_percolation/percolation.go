package compute_percolation

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/sbezverk/algorithms/pkg/quick_union_weighted_pathcompression"
)

type Percolation interface {
	Open(int, int) error
	IsOpen(int, int) (bool, error)
	IsFull(int, int) (bool, error)
	NumberOfOpenedSites() int
	Percolates() bool
}

var _ Percolation = &percolation{}

type percolation struct {
	gridSize     int
	nOpenedSites int
	qu           quick_union_weighted_pathcompression.QUWPC
}

func (p *percolation) NumberOfOpenedSites() int {
	return p.nOpenedSites
}

func (p *percolation) getOpenedNeighbors(row, col int) []int {
	neighbors := make([]int, 0)
	for r := row - 1; r <= row+1; r++ {
		if r < 0 {
			continue
		}
		if r*p.gridSize > p.gridSize*p.gridSize {
			break
		}
		for c := col - 1; c <= col+1; c++ {
			if c-1 > 0 {
				// check for neighbor at top left corner
				if opened, _ := p.qu.Opened(r*p.gridSize + c - 1); opened {
					neighbors = append(neighbors, r*p.gridSize+c-1)
				}
			}
			if opened, _ := p.qu.Opened(r*p.gridSize + c); opened {
				neighbors = append(neighbors, r*p.gridSize+c)
			}
			if c+1 <= p.gridSize {
				// check for neighbor at top right corner
				if opened, _ := p.qu.Opened(r*p.gridSize + c + 1); opened {
					neighbors = append(neighbors, r*p.gridSize+c+1)
				}
			}

		}
	}

	return neighbors
}

func (p *percolation) Open(row, col int) error {
	cSite := p.gridSize*row + col
	if err := p.qu.Open(cSite); err != nil {
		return err
	}
	// New site is opened
	p.nOpenedSites++
	// Need to establish union with all Opened neighbors in the grid
	for _, n := range p.getOpenedNeighbors(row, col) {
		if err := p.qu.Union(cSite, n); err != nil {
			return err
		}
	}

	return nil
}

func (p *percolation) IsOpen(row, col int) (bool, error) {
	cSite := p.gridSize*row + col

	return p.qu.Opened(cSite)
}

func (p *percolation) IsFull(row, col int) (bool, error) {
	cSite := p.gridSize*row + col

	return p.qu.Opened(cSite)
}

func (p *percolation) Percolates() bool {
	// Check if first virtual site, element 0 is connected to the second virtual site
	// last element of the grid.
	r, _ := p.qu.Connected(0, p.gridSize*p.gridSize+1)
	return r
}

func NewPercolation(n int) (Percolation, error) {
	p := &percolation{
		gridSize: n,
	}
	// Creating Grid n*n + 2 virtual sites, 1st virtual site is connected to all top row elements,
	// 2nd virtual site is connected to all bottom elements. Percolation is checked between these 2 virtual nodes.
	qu, err := quick_union_weighted_pathcompression.NewQUWPC(n*n + 2)
	if err != nil {
		return nil, err
	}
	p.qu = qu
	// Connecting virtual sites to first and last rows of the grid
	for i := 1; i <= n; i++ {
		if err := p.qu.Union(0, i); err != nil {
			return nil, err
		}
		if err := p.qu.Union(n*n+1, n*n+1-i); err != nil {
			return nil, err
		}
	}

	return p, nil
}

type PercolationStats interface {
	Mean() float64
	StdDev() float64
	Confidence() (float64, float64)
}

var _ PercolationStats = &percolationStats{}

type percolationStats struct {
	probabilities []float64
	mean          float64
	stddev        float64
	confidenceLo  float64
	confidenceHi  float64
}

func (ps *percolationStats) Mean() float64 {
	ps.mean = 0
	for i := 0; i < len(ps.probabilities); i++ {
		ps.mean += float64(ps.probabilities[i])
	}
	return ps.mean / float64(len(ps.probabilities))
}

func (ps *percolationStats) StdDev() float64 {
	m := ps.Mean()
	for i := 0; i < len(ps.probabilities); i++ {
		ps.stddev += math.Pow((float64(ps.probabilities[i]) - m), 2)
	}

	return math.Sqrt(ps.stddev / float64(len(ps.probabilities)))
}

func (ps *percolationStats) Confidence() (float64, float64) {
	m := ps.Mean()
	s := ps.StdDev()
	ps.confidenceLo = m - 1.96*s/math.Sqrt(float64(len(ps.probabilities)))
	ps.confidenceHi = m + .96*s/math.Sqrt(float64(len(ps.probabilities)))
	return ps.confidenceLo, ps.confidenceHi
}

func ComputePercolationStats(gridSize int, trials int) (PercolationStats, error) {
	ps := &percolationStats{
		probabilities: make([]float64, trials),
	}
	for trial := 0; trial < trials; trial++ {
		p, err := NewPercolation(gridSize)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize percolation with error: %+v", err)
		}

		src := rand.New(rand.NewSource(time.Now().UnixNano()))
		percolates := false
		done := false
		for !done {
			if p.NumberOfOpenedSites() == gridSize*gridSize {
				done = true
				continue
			}
			s := src.Intn(gridSize*gridSize + 1)
			if s == 0 {
				// Ignoring first element, which is a virtual site
				continue
			}
			row := s / gridSize
			col := s - (row * gridSize)
			opened, err := p.IsOpen(row, col)
			if err != nil {
				return nil, fmt.Errorf("failed to check if a site row: %d col: %d is opened, with error: %+v", row+1, col, err)
			}
			if opened {
				continue
			}
			if err := p.Open(row, col); err != nil {
				return nil, fmt.Errorf("failed to open site row: %d col: %d with error: %+v", row+1, col, err)

			}
			if p.Percolates() {
				percolates = true
				done = true
			}
		}
		if percolates {
			// fmt.Printf("Percolation is found with number of opened sites: %d\n", p.NumberOfOpenedSites())
			ps.probabilities[trial] = float64(p.NumberOfOpenedSites()) / float64(gridSize*gridSize)
		} else {
			fmt.Printf("Percolation is not found for trial %d\n", trial)
			ps.probabilities[trial] = 1 / float64(gridSize*gridSize)
		}
	}

	return ps, nil
}
