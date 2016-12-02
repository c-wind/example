package main

import (
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("hello\n"))
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	if err := http.ListenAndServe(":9000", nil); err != nil {
        panic(err.Error())
    }
}
