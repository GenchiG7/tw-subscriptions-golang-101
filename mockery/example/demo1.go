package demo1

type Person struct {
	Name string
}

//mockery --name Mydemo1
type Mydemo1 interface {
	DemoFunc1() string
	DemoFunc2(p Person) string
	DemoFunc3(p *Person) string
	DemoFunc4(p *Person) (string, error)
}
