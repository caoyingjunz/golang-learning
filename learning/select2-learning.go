package main

import (
	"fmt"
	"time"
)

type TestA interface {
	Print(s string)
}

type Test struct{}

func (t *Test) Print(s string) {
	fmt.Println("test")
}

func main() {

	t := TestA(&Test{})
	t.Print("tt")

	ready := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second)
		close(ready)
	}()

	// Log error every connectionLoggingInterval
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// Wait until Dial() succeeds.
	for {
		select {
		case <-ticker.C:
			fmt.Println("Still connecting to test")

		case <-ready:
			fmt.Println("ready")
			return
		}
	}

}
