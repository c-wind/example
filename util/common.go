package util

import (
    "fmt"
    "time"
)

func Log(s string) {
    fmt.Printf("%v %v\n", time.Now(), s)
}
