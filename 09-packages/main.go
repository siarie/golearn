package main

import "fmt"
import "golearn/09-packages/math"

func main() {
	xs := []float64{2, 1, 6, 4, 8, 6}

	avg := math.Average(xs)
	min := math.Min(xs)
	max := math.Max(xs)
	fmt.Println(avg, min, max)
}
