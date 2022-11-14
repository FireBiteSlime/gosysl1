package t3

import "fmt"

func loadArr(arr [5]int, c chan<- int) {
	for _, value := range arr {
		c <- value
	}
	close(c)
}

func pow(req <-chan int, powres chan<- int) {
	for value := range req {
		powres <- (value * value)
	}
	close(powres)
}

func powSumm(powres <-chan int, summres chan<- int) {
	var res int
	for value := range powres {
		res += value
	}
	summres <- res
	close(summres)

}

func out(summres <-chan int) {
	for value := range summres {
		fmt.Println("Сумма = ", value)
	}
}

func Test() {
	var arr = [5]int{2, 4, 6, 8, 10}
	req := make(chan int)
	powres := make(chan int)
	summres := make(chan int)

	go loadArr(arr, req)
	go pow(req, powres)
	go powSumm(powres, summres)
	out(summres)

}
