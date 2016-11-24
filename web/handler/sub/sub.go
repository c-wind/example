package sub

import (
	"fmt"
	"net/http"
)

func sub(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, "sub")
}
