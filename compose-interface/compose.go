package main

type A interface {
	hello() string
}

type B interface {
	hello() string
}

type C interface {
	A
	B
}

func main() {
	var ob C = &work{}
	ob.hello()
}

type work struct{}

func (w *work) hello() string {
	return "hello"
}
