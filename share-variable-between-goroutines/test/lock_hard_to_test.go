package test

import (
	"sync"
	"testing"
)

func SomeBusinessFunc() {
	flag := true
	var l sync.Mutex
	l.Lock()
	go func1(&flag, &l)

	func2(&flag, &l)
	for {
		continue
	}
}

func func1(flag *bool, l *sync.Mutex) {
	// other business

	l.Lock()
	for *flag {
		continue
	}

	// other business
}
func TestFunc1(t *testing.T) {

}

func func2(flag *bool, l *sync.Mutex) {
	// other business
	*flag = false
	l.Unlock()
}
func TestFunc2(t *testing.T) {

}
