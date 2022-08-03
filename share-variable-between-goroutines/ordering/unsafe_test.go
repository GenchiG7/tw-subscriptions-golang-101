package testordering

import (
	"fmt"
	"sync"
	"testing"
)

func TestUnsafe(t *testing.T) {
	i := 0

	for {
		i++
		var a, b int
		//fmt.Printf("111 %d, %d\n", a, b)
		wg := new(sync.WaitGroup)
		wg.Add(1)
		go func() {
			defer wg.Done()
			a = 2
			b = 1
		}()

		if b == 1 && a == 0 {
			fmt.Printf("%dth round, Non thread safe!", i)
			break
		}

		wg.Wait()
		//fmt.Printf("222 %d, %d\n", a, b)
	}
}
