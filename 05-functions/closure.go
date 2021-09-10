package main

import "fmt"

func main() {
	// add number
	add := func(num ...int) (total int) {
		total = 0

		for _, v := range num {
			total += v
		}

		return
	}

	fmt.Println(add(1, 2, 3, 4, 5, 6, 7))

	nextEven := evenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}

func evenGenerator() func() uint {
	i := uint(0)

	return func() uint {
		res := i

		i += 2
		return res
	}
}
