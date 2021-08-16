package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Println("03. Flow Control - Loop For")
	fmt.Println("===========================")
	fmt.Println("")

	// basic loop syntax
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println("Sum of 1..100 is", sum)

	// Empty prestatements and poststatements
	// Similar to `while`
	var num int64
	for num != 6 {
		num = rand.Int63n(15)
		fmt.Printf("%d ", num)
	}

	fmt.Println("")

	// Infinite loop, break and continue
	var randNum int64
	for {
		fmt.Println("Infinite Loop...")
		if randNum = rand.Int63n(10); randNum == 5 {
			fmt.Println("And finish until meet the requirement")
			break
		}
		fmt.Println(randNum)
	}

	// continue:
	// sum even number
	sumEven := 0
	for j := 1; j <= 100; j++ {
		if j%2 != 0 {
			continue
		}
		sumEven += j
	}

	fmt.Println("Sum of even number 1..100.", sumEven)
}
