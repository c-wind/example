package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"strings"
)

const BufferSize = 1024

// 接受一个TCPConn处理内容
func handleConn(id int, conn *net.TCPConn) {
	defer func() {
		conn.Close()
		if err := recover(); err != nil {
			fmt.Printf("%d panic:%v\n", id, err)
		}
	}()

	buf := make([]byte, BufferSize)

	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Printf("%d %s closed!\n", id, conn.RemoteAddr().String())
			return
		}

		if err != nil {
			panic(err.Error())
		}

		msg := strings.TrimSpace(string(buf[:n]))

		if msg == "exit" {
			fmt.Printf("%d %s exit\n", id, conn.RemoteAddr().String())
			return
		}

		fmt.Printf("%d recv:%s, from:%s\n", id, msg, conn.RemoteAddr().String())

		conn.Write(buf[:n])
	}
}

func echoServer(host string) {
	addr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		panic(err.Error())
	}

	// 开始监听
	ln, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		panic(err.Error())
	}
	defer ln.Close()

	for i := 0; ; i++ {
		conn, err := ln.AcceptTCP()
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("new client[%d]:%s has connected!\n", i, conn.RemoteAddr().String())
		go handleConn(i, conn) //起一个goroutine处理
	}

}

func main() {
	// 参数，默认绑定9000端口
	port := flag.Int("p", 9000, "listen port")
	flag.Parse()

	host := fmt.Sprintf(":%d", *port)

	echoServer(host)
}
