package example

import (
	"fmt"
)

type simpleError struct {
	code int
	msg  string
}

func DoPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err:%+v\n", err)
		}
	}()

	panic(simpleError{-1, "not found"})
}
