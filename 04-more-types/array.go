package main

import "fmt"

func main() {
	// Array
	var x [5]int

	x[2] = 30

	fmt.Println(x)
	fmt.Println("The value of array at index 2 is", x[2])

	var grade [5]float64
	grade[0] = 12
	grade[1] = 21
	grade[2] = 11
	grade[3] = 77
	grade[4] = 30

	gradeLen := len(grade)

	var average float64 = 0
	// for i := 0; i < gradeLen; i++ {
	// 	average += grade[i]
	// }

	for _, score := range grade {
		average += score
	}

	fmt.Println("Average:", average/float64(gradeLen))

	// short syntax array declaration
	arr := [3]float64{12, 34, 65}
	fmt.Println(arr)
}
