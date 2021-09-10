/**
 * Defer Panic and Recover
 */
package main

import "fmt"

func defered() {
	fmt.Println("defer will be run after the function completes")
}

func main() {
	defer defered()

	fmt.Println("Defer, Panic & Recover")

	defer func() {
		// recover stops the panic and returns the value that was passed to the call to panic
		str := recover()
		fmt.Println(str)
	}()

	panic("Duarrr...")
}
