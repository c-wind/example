package main

import (
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func work(id int) {
	println(id, ", begin:", time.Now().Format(time.ANSIC))
	time.Sleep(time.Second)
	wg.Done()
	println(id, ",   end:", time.Now().Format(time.ANSIC))
}

func main() {
	println("main begin:", time.Now().Format(time.ANSIC))
	go work(1)
	go work(2)
	go work(3)
	go work(4)
	go work(5)

	wg.Add(5)
	wg.Wait()

	println("main end:", time.Now().Format(time.ANSIC))
}
