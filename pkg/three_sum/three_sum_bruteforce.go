package three_sum

func ThreeSumBruteForce(a []int) int {
	count := 0
	l := len(a)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if a[i]+a[j]+a[k] == 0 {
					count++
				}
			}
		}
	}

	return count
}
