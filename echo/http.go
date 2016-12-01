package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"runtime"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	msg := req.URL.Query().Get("msg")
	//    fmt.Printf("recv msg:%v\n", msg)
	io.WriteString(w, msg+"\n")
}

func main() {
	// 参数，默认绑定9000端口
	port := flag.String("p", "9000", "listen port")
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
