package test

import "testing"

func CSP() {
	done := newfunc2()
	newfunc1(done)
}

func newfunc1(done <-chan interface{}) {
	go func() {
		// other business
		select {
		case <-done:
			break
		}

		// other business
	}()
}
func TestNewfunc1(t *testing.T) {

}

func newfunc2() <-chan interface{} {
	done := make(chan interface{})
	go func() {
		// other business

		close(done)
	}()
	return done
}
func TestNewfunc2(t *testing.T) {

}
