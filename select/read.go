package main

import (
	"time"
)

func main() {
	c := make(chan bool, 1)
	timeout := time.After(time.Second)

	//3 秒后向chan写数据
	go func() {
		time.Sleep(time.Second * 3)
		c <- true
	}()

	select {
	case <-c:
		println("recv")
	case <-timeout:
		println("timeout!")
	}
}
