package main

import (
	"fmt"
)

func main() {
	defer func() {
		println("1")
	}()
	defer func() {
		println("2")
		fmt.Printf("%v\n", recover())
	}()
	defer func() {
		println("3")
	}()

	panic("i am panic")
}
