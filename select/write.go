package main

import (
	"time"
)

func main() {
	c := make(chan bool, 1)

	//每3秒后读chan数据
	go func() {
		time.Sleep(time.Second * 3)
		v := <-c
		println("recv ", v)
	}()

	//第一次写，会成功, 因为有缓冲1
	select {
	case c <- true:
		println("send ok")
	default:
		println("full!")
	}

	//此时chan已满，因为还没有读出数据
	select {
	case c <- true:
		println("send ok")
	default:
		println("full!")
	}
}
