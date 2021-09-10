package main

import "fmt"

// create interface
type Person interface {
	greet() string
}

type Human struct {
	Name string
}

func (h *Human) greet() string {
	return "Hi, I'am " + h.Name
}

func main() {
	fmt.Println("7.2. Interfaces")
	fmt.Println("===============")

	var a Human = Human{"Michael"}

	fmt.Println(a.greet())
}
