package main

import "fmt"

/*
  main function
*/
func main() {
	// call other function
	f()

	var a int = 12
	var b int = 60
	fmt.Printf("%d * %d = %d\n", a, b, multiply(a, b))

	x, y := multiValue()
	fmt.Println(x, y)

	fmt.Println(sayHello())
}

/*
   void function without params
*/
func f() {
	fmt.Println("05. Functions")
}

/*
   function with params and returning value
*/
func multiply(a int, b int) int {
	return a * b
}

/*
   Return multiple value
*/
func multiValue() (int, float64) {
	return 12, 90.12
}

/*
   naming return type
*/
func sayHello() (greeting string) {
	greeting = "Hello, world!"
	return
}
