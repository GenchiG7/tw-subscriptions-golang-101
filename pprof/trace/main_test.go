package tracetest

import (
	"testing"
)

func consume(in <-chan interface{}) <-chan interface{} {
	finishe := make(chan interface{})
	go func() {
		defer close(finishe)
		select {
		case <-in:
		}
	}()
	return finishe
}

func produce(out chan<- interface{}) {
	var sum uint64
	for i := uint64(0); i < 100; i++ {
		sum += i * i
	}
	out <- sum
}

var concurrencyNum int = 5

func TestProducerConsumer(t *testing.T) {
	buffer := make(chan interface{}, concurrencyNum)
	finishChans := make([]<-chan interface{}, concurrencyNum)

	for i := 0; i < concurrencyNum; i++ {
		produce(buffer)
	}

	for i := 0; i < concurrencyNum; i++ {
		finishChans[i] = consume(buffer)
	}

	for i := 0; i < concurrencyNum; i++ {
		<-finishChans[i]
	}
}
