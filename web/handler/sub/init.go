package sub

import (
	"github.com/c-wind/example/web/handler"
)

func init() {
	handler.Router.Path("/sub/{operation}").Methods("POST").HandlerFunc(sub)
}
