package testvisibilit

import (
	"fmt"
	"sync"
	"testing"
)

func TestSafe(t *testing.T) {
	flag := true
	var l sync.Mutex
	l.Lock()
	go func() {
		l.Lock()
		for flag {
			continue
		}
		fmt.Printf("Never end\n")
	}()

	flag = false
	l.Unlock()
	for {
		continue
	}
}
