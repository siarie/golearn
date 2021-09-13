package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)

	if xslen := len(xs); xslen == 0 {
		return 0
	}

	for _, x := range xs {
		total += x
	}

	return total / float64(len(xs))
}

// Find the minimum value of a series of numbers
func Min(xs []float64) float64 {
	var min float64 = xs[0]

	for _, x := range xs {
		if x < min {
			min = x
		}
	}

	return min
}

// Find the maximum value of a series of numbers
func Max(xs []float64) float64 {
	var max float64 = xs[0]

	for _, x := range xs {
		if x > max {
			max = x
		}
	}

	return max
}
