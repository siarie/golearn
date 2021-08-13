package main

import "fmt"

/**
 * Learn Golang 02
 *
 * Variable, Basic Data Type, Aritmetic Operator and
 * Type conversion
 */
func main() {
	// variable declaration
	var name string
	// variable assignment
	name = "Sri Aspari"
	fmt.Printf("Name:\t%s\n", name)

	name = "siarie" // change value
	fmt.Printf("Nick:\t%s\n", name)
	fmt.Printf("=====================\n\n")

	// multiple variable declaration
	var x, y int
	x, y = 23, 8

	// Aritmetic Operator
	fmt.Println("1. Basic math operations:")
	fmt.Printf("   %d + %d = %d\n", x, y, x+y) // addition
	fmt.Printf("   %d - %d = %d\n", x, y, x-y) // subtraction
	fmt.Printf("   %d * %d = %d\n", x, y, x*y) // multiplication
	fmt.Printf("   %d / %d = %d\n", x, y, x/y) // division
	fmt.Printf("   %d %% %d = %d\n", x, y, x%y) // modulo division
	fmt.Println("")

	fmt.Println("2. Volume of Circle:")
	// support unicode character
	var π float64 = 3.14 // declare and assign
	r := 7 // type inference

	fmt.Printf("   Given a circle with r = %d, and π = %.2f.\n", r, π)
	fmt.Println("   What is the volume of the circle?...")

	// solution
	d := float64(r*r)
	V := d * π
	fmt.Println("   Soln:")
	fmt.Printf("   \tV = π * r^2\n")
	fmt.Printf("   \tV = %d * %.2f\n", r, π)
	fmt.Printf("   \tV = %.2f\n", V)
	fmt.Printf("   The volume of circle is %.2f\n", V)
	fmt.Println("")
	fmt.Printf(" — signature: %s\n\n", name)
}
