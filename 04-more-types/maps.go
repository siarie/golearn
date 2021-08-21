package main

import "fmt"

func main() {
	grade := make(map[string]int)
	grade["math"] = 8
	grade["biology"] = 9

	fmt.Println(grade)
	fmt.Println(grade["math"])

	// delete key
	delete(grade, "math")
	fmt.Println(grade)

	// check is key exists
	name, ok := grade["biology"]
	fmt.Println(name, ok)

	// shorter syntax to create array
	allGrade := map[string]float64{
		"smt1": 200,
		"smt2": 210,
		"smt3": 190,
	}

	fmt.Println(allGrade)
}
