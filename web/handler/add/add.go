package add

import (
	"fmt"
	"net/http"
)

func add(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, "add")
}
