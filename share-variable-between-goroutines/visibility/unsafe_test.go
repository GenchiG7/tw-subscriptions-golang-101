package testvisibilit

import (
	"fmt"
	"testing"
)

func TestUnsafe(t *testing.T) {
	flag := true

	go func() {
		for flag {
			continue
		}
		fmt.Printf("Never end\n")
	}()

	flag = false

	for {
		continue
	}
}
