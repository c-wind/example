package main

import (
	"github.com/c-wind/example/channel"
	"github.com/c-wind/example/util"
)

func main() {
	util.Log("test")

	c := make(chan int, 10)
	go channel.DoSend(c)
	sum := channel.DoRecv(c)
	println(sum)

	channel.DoSend(c)

}
