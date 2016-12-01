package main

import (
	"os"
	"sync"
	"time"
)

var (
	last time.Time
	mu   sync.Mutex
)

func modifyFile() error {
	mu.Lock()
	f, err := os.Open("/tmp/abc.txt")
	if err != nil {
		mu.Unlock()
		return err
	}

	if _, err = f.WriteString("ok"); err != nil {
		mu.Unlock()
		f.Close()
		return err
	}

	last = time.Now()
	mu.Unlock()

	return nil
}

func modifyFileWithDefer() error {
	mu.Lock()
	defer mu.Unlock()

	f, err := os.Open("/tmp/abc.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString("ok"); err != nil {
		return err
	}

	last = time.Now()
	return
}

func main() {
	modifyFile()
	modifyFileWithDefer()
}
