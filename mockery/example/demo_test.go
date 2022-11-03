package demo1_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

	demo1 "mockerydemo/example"
	"mockerydemo/example/mocks"
)

// simple example
func TestDemo1(t *testing.T) {
	demo := mocks.NewMydemo1(t)

	demo.On("DemoFunc1").Return("hello word")

	fmt.Printf("%s\n", demo.DemoFunc1())
}

// repeat
func TestDemo2(t *testing.T) {
	demo := mocks.NewMydemo1(t)

	demo.On("DemoFunc1").Return("hello word 1")

	fmt.Printf("%s\n", demo.DemoFunc1())
	fmt.Printf("%s\n", demo.DemoFunc1())
}

// valid numbers of call
func TestDemo3(t *testing.T) {
	demo := mocks.NewMydemo1(t)

	demo.On("DemoFunc1").Return("hello word 1").Once()

	fmt.Printf("%s\n", demo.DemoFunc1())
	fmt.Printf("%s\n", demo.DemoFunc1())
}

// return different result
func TestDemo4(t *testing.T) {
	demo := mocks.NewMydemo1(t)

	demo.On("DemoFunc1").Return("hello word 1").Maybe().Once()
	demo.On("DemoFunc1").Return("hello word 2").Once()

	fmt.Printf("%s\n", demo.DemoFunc1())
	fmt.Printf("%s\n", demo.DemoFunc1())
}

// verify input parameter: happy
func TestDemo5(t *testing.T) {
	demo := mocks.NewMydemo1(t)

	p := demo1.Person{"G7"}

	demo.On("DemoFunc2", p).Return(p.Name)

	fmt.Printf("%s\n", demo.DemoFunc2(p))
}

// verify input parameter: unhappy
func TestDemo6(t *testing.T) {
	demo := mocks.NewMydemo1(t)

	p1 := demo1.Person{"G7"}
	p2 := demo1.Person{"Genchi"}

	demo.On("DemoFunc2", p1).Return(p1.Name)

	fmt.Printf("%s\n", demo.DemoFunc2(p2))
}

// verify input parameter: different obj with same value
func TestDemo7(t *testing.T) {
	demo := mocks.NewMydemo1(t)

	p1 := demo1.Person{"G7"}
	p2 := demo1.Person{"G7"}

	demo.On("DemoFunc2", p1).Return(p1.Name)

	fmt.Printf("%s\n", demo.DemoFunc2(p2))
}

// verify ptr
func TestDemo8(t *testing.T) {
	demo := mocks.NewMydemo1(t)

	p := &demo1.Person{"G7"}

	demo.On("DemoFunc3", p).Return("hello world")

	fmt.Printf("%s\n", demo.DemoFunc3(p))
}

// verify ptr: different obj with same value
func TestDemo9(t *testing.T) {
	demo := mocks.NewMydemo1(t)

	p1 := &demo1.Person{"G7"}
	p2 := &demo1.Person{"G7"}

	demo.On("DemoFunc3", p1).Return("hello world")

	fmt.Printf("%s\n", demo.DemoFunc3(p2))
}

// without verify input
func TestDemo10(t *testing.T) {
	demo := mocks.NewMydemo1(t)

	p := &demo1.Person{"G7"}

	demo.On("DemoFunc3", mock.Anything).Return("hello world")

	fmt.Printf("%s\n", demo.DemoFunc3(p))
}

// without verify input: bad smells of any, should fail but pass
func TestDemo11(t *testing.T) {
	demo := mocks.NewMydemo1(t)

	demo.On("DemoFunc3", mock.Anything).Return("hello world")

	fmt.Printf("%s\n", demo.DemoFunc3(nil))
}

//new feature: --with-expecter
func TestDemo12(t *testing.T) {
	demo := mocks.NewMydemo1(t)

	p := &demo1.Person{"G7"}
	// it's not correct
	//demo.On("DemoFunc4", mock.Anything, mock.Anything, mock.Anything).Return("hello world")

	demo.EXPECT().DemoFunc4(p).Return(errors.New("some error"), "hello world")

	res, _ := demo.DemoFunc4(p)
	fmt.Printf("%s\n", res)
}
