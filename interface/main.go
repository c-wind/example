package main

type Person interface {
	Hi()
}

type Man struct {
}

func (m *Man) Hi() {
	println("Man")
}

type Woman struct {
}

func (w *Woman) Hi() {
	println("Woman")
}

func start(p Person) {
	p.Hi()
}

func main() {
	start(Man{})
	start(&Man{})
	start(&Woman{})
}
