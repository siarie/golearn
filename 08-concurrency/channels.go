package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping!"
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong!"
	}
}

func printer(c <-chan string) {
	for {
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}

/* Sleep function using `time.After()`
 */
func Sleep(s int) {
	<-time.After(time.Second * time.Duration(s))
}

func main() {
	fmt.Println("08. Concurrency")
	Sleep(8)
	// var c chan string = make(chan string)

	// go pinger(c)
	// go ponger(c)
	// go printer(c)

	// Select
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan int, 20)

	go func() {
		for {
			c1 <- "channel 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "channel 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {

		for i := 100; ; i++ {
			c3 <- i
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case msg3 := <-c3:
				fmt.Println(msg3)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
