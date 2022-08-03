package testordering

import (
	"fmt"
	"sync"
	"testing"
)

func TestSafe(t *testing.T) {
	i := 0

	for {
		i++
		var a, b int
		wg := new(sync.WaitGroup)
		wg.Add(1)
		var l sync.Mutex
		l.Lock()
		go func() {
			defer wg.Done()
			a = 2
			b = 1
			l.Unlock()
		}()

		l.Lock()
		if b == 1 && a == 0 {
			fmt.Printf("%dth round, Non thread safe!", i)
			break
		}

		wg.Wait()
	}
}
