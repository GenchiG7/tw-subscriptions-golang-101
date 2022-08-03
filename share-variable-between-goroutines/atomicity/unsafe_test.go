package testatomicity

import (
	"fmt"
	"sync"
	"testing"
)

func TestUnsafe(t *testing.T) {
	concurrencyNum := 100

	wg := new(sync.WaitGroup)
	wg.Add(concurrencyNum)

	var share int64
	for i := 0; i < concurrencyNum; i++ {
		go func(a *int64, wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				*a++
			}
		}(&share, wg)

	}

	wg.Wait()

	fmt.Printf("share is %d\n", share)
}
