package main

import "fmt"

func main() {
	grade := []int{7, 8, 6, 9, 7, 7, 8, 9, 6, 8}
	gradeSmt1 := grade[0:5]

	fmt.Println(grade)
	fmt.Println(gradeSmt1)

	// append
	grade2 := append(grade, 10)
	fmt.Println(grade2)

	// copy
	var gradeSlice = make([]int, 6)
	gradeCopy := copy(gradeSlice, gradeSmt1)
	fmt.Println(gradeCopy)

	// create empty slice
	// make(arr, min, max)
	emptySlice := make([]int, 5, 10)
	fmt.Println(emptySlice)
}
