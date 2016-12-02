package main

const (
	A1 = iota
	A2
	A3
)

const (
	B1 = iota
	B2
	B3
)

const (
	C0 = iota
	C1 = 1 << iota
	C2
	C3
)

const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {

	println(A1, A2, A3)
	println(B1, B2, B3)

	println(int(KB), int(MB))
	println(C0, C1, C2, C3)
}
