package main

import (
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		c1 <- 1
		time.Sleep(time.Second * 3)
		c2 <- 2
		for {
		}
	}()

	run := true

	for run {
		select {
		case i := <-c1:
			println("c1", i)
		case i := <-c2:
			println("c2", i)
		default:
			println("default", len(c1), len(c2))
		}
	}

}
