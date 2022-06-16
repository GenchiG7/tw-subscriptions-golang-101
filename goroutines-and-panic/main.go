package main

import (
	"fmt"
	"time"
)


func main() {

	// start a background goroutine
	go func(){
		
		// background goroutine start another goroutine every 3s
		for {
			go panicFunc()
			time.Sleep(3 * time.Second)
		}
	}()

	for {
		fmt.Printf("main goroutine is alive...\n")
		time.Sleep(5 * time.Second)
	}
}

func panicFunc() {
	panic("something wrong")
}