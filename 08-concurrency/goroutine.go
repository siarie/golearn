package main

import (
	"fmt"
	"math/rand"
	"time"
)

func loop(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
		amt := time.Duration(rand.Intn(2000))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	// goroutine
	go loop(100)

	var input string
	fmt.Scanln(&input)

	fmt.Println("=======")
	fmt.Println(input)
}
