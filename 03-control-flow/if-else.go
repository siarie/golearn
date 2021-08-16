package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("03. Flow Control")
	fmt.Printf("================\n\n")

	rand.Seed(time.Now().UnixNano())

	var a, b, c int
	a, b = rand.Intn(50-1)+1, rand.Intn(50-1)+1

	fmt.Printf("What is %d + %d? ", a, b)
	fmt.Scanln(&c)

	fmt.Printf("\nYour answer %d\n", c)

	if correctAnswer := a + b; c == correctAnswer {
		fmt.Println("You're right.")

		fmt.Printf("\nAnswer with \"yes\" or \"no\".")
		fmt.Printf("\nIs %d even? ", c)
		var answer string
		fmt.Scanln(&answer)

		// nested if
		if answer == "yes" && c%2 == 0 {
			fmt.Println("Wew.. you're smart.")
		} else if answer == "no" && c%2 != 0 { // else if
			fmt.Println("^_^ Correct..")
		} else {
			fmt.Println("-_- Uhh.. that wrong answer.")
		}
	} else {
		fmt.Println("")
		fmt.Println("You're wrong.")
		fmt.Printf("Correct Answer: %d\n", correctAnswer)
	}
}
