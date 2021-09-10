package main

import "fmt"

func greeting(textPtr *string) {
	*textPtr = "Hello,"
}

func username(namePtr *string) {
	*namePtr = "siarie"
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x, y *int) {
	var tmp = *x
	*x = *y
	*y = tmp
}

func main() {
	fmt.Println("06. Pointer")

	greet := "Welcome,"
	greeting(&greet)

	// get pinter using `new`
	uname := new(string)
	username(uname)

	fmt.Println(greet, *uname)

	a := 5.0
	square(&a)
	fmt.Println(a)

	// swap two number
	x := 1
	y := 2

	fmt.Printf("before swap: x=%d, y=%d\n", x, y)
	swap(&x, &y)

	fmt.Printf(" after swap: x=%d, y=%d\n", x, y)
}
