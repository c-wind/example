package main

import (
	"runtime"
	"time"
)

func test() {
	time.Sleep(time.Minute)
}

func main() {
	println("111111111111")
	runtime.GOMAXPROCS(runtime.NumCPU())
	println("222222222222")
	go test()
	go test()
	go test()
	go test()
	go test()

	time.Sleep(time.Minute)
}
