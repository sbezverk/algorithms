package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sbezverk/algorithms/pkg/compute_percolation"
)

var (
	gridSize int
	trials   int
)

func init() {
	flag.IntVar(&gridSize, "grid-size", 100, "the size of the grid, default 100x100")
	flag.IntVar(&trials, "trials", 100, "number of percolation computation trials, default 100")
}

func main() {
	flag.Parse()
	ps, err := compute_percolation.ComputePercolationStats(gridSize, trials)
	if err != nil {
		fmt.Printf("Computation of percolation failed with error: %+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Percolation stats for grid of %d x %d and for %d trials:\n", gridSize, gridSize, trials)
	fmt.Printf("\t\tMean: %f\n", ps.Mean())
	fmt.Printf("\t\tStd Dev: %f\n", ps.StdDev())
	cl, ch := ps.Confidence()
	fmt.Printf("\t\tConfidence: [ %f, %f ]\n", cl, ch)
}
