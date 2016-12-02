package main

func main() {
	str := []byte("你好，测试!")

	for i, v := range str {
		println(i, string(v))
	}

	println("-----------------")

	for i, v := range string(str) {
		println(i, string(v))
	}

}
