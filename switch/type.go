package main

func typeCheck(i interface{}) {
	switch o := i.(type) {
	case int:
		println("int", o)
	case string:
		println("string", o)
	}
}

func main() {
	typeCheck(1)
	typeCheck("ABC")
}
