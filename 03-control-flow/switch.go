package main

import "fmt"

func main() {
	fmt.Println("03. Control Flow - Switch")
	fmt.Println("")

	fmt.Println("Input number between 1 and 12.")
	fmt.Printf("> ")

	var month int
	fmt.Scanln(&month)

	// basic switch syntax
	switch month {
	case 1:
		fmt.Printf("%d -> January\n", month)
	case 2:
		fmt.Printf("%d -> February\n", month)
	case 3:
		fmt.Printf("%d -> March\n", month)
	case 4:
		fmt.Printf("%d -> April\n", month)
	case 5:
		fmt.Printf("%d -> May\n", month)
	case 6:
		fmt.Printf("%d -> June\n", month)
	case 7:
		fmt.Printf("%d -> July\n", month)
	case 8:
		fmt.Printf("%d -> August\n", month)
	case 9:
		fmt.Printf("%d -> September\n", month)
	case 10:
		fmt.Printf("%d -> October\n", month)
	case 11:
		fmt.Printf("%d -> November\n", month)
	case 12:
		fmt.Printf("%d -> December\n", month)
	default:
		fmt.Println("Only number 1-12 allowed.")
	}

	// multiple expressions
	switch month {
	case 1, 2, 3, 4, 5, 6:
		fmt.Println("First Semester.")
	case 7, 8, 9, 10, 11, 12:
		fmt.Println("Second Semester.")
	}

	// omit a condition
	switch {
	case month <= 6:
		fmt.Println("It's Dry")
	case month <= 12:
		fmt.Println("It's Rainy")
	case month > 12:
		fmt.Println("???")
	}
}
