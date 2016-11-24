package channel

// DoSend send value.
func DoSend(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}

	close(c)
}

//DoRecv recv value.
func DoRecv(c chan int) int {
	var sum int
	for v := range c {
		sum += v
	}
	return sum
}
