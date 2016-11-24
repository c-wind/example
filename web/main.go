package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/c-wind/example/web/handler"
	_ "github.com/c-wind/example/web/handler/add"
	_ "github.com/c-wind/example/web/handler/sub"
)


func main() {
	defer func() {
		if panicRecover := recover(); panicRecover != nil {
			fmt.Printf("Main func failed for %v", panicRecover)
		}
	}()

	runtime.GOMAXPROCS(runtime.NumCPU())

	http.Handle("/", handler.Router)

	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Println("Start handler server error:", err)
		panic(err)
	}
}
