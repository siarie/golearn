package main

import "fmt"

func main() {

	fmt.Println("John's average score is", average(7, 8, 9, 6, 6, 9))

	scoreKelly := []int{6, 6, 8, 7, 9, 8}
	fmt.Println("Kelly's average score is", average(scoreKelly...))
}

/*
   Calculate average using variadic function
*/
func average(grades ...int) float64 {
	total := 0

	for _, v := range grades {
		total += v
	}

	average := float64(total) / float64(len(grades))

	return average
}
