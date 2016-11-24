package handler

import (
    "net/http"

    "github.com/gorilla/mux"
)

var (
    Router = mux.NewRouter()
)

const (
    URL_VERSION = "v1"
)

func init() {
    Router.NotFoundHandler = http.NotFoundHandler()
}
