package main

type Person struct{
}

type Man struct {
    Person
}

func (p *Person) Hello() {
    println("hello Person")
}

func (m *Man) Hello() {
    println("Man")
}

type Woman struct {
    Person
}

func (w *Woman) Hello() {
    println("Woman")
}


func main() {
    m := Man{}
    m.Hello()

    w := Woman{}
    w.Hello()

    m.Person.Hello()
}
