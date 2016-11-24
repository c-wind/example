package add

import (
	"github.com/c-wind/example/web/handler"
)

func init() {
	handler.Router.Path("/add/{operation}").Methods("POST").HandlerFunc(add)
}
