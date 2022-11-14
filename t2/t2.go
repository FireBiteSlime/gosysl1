package t2

import "fmt"

func loadArr(arr [5]int, c chan<- int) {
	for _, value := range arr {
		c <- value
	}
	close(c)
}

func pow(req <-chan int, res chan<- int) {
	for value := range req {
		res <- (value * value)
	}
	close(res)
}

func out(res <-chan int) {
	for value := range res {
		fmt.Println(value)
	}
}

func Test() {
	var arr = [5]int{2, 4, 6, 8, 10}
	req := make(chan int)
	res := make(chan int)

	go loadArr(arr, req)
	go pow(req, res)
	out(res)
}
