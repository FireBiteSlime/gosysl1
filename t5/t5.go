package t5

import (
	"fmt"
	"time"
)

func launch(main chan int, tmr chan int) {
	var seconds int
	fmt.Println("Введите колличество секунд: ")
	fmt.Scanf("%d\n", &seconds)

	go timer(tmr, seconds)
	go reader(main)

	for i := 0; ; i++ {
		select {
		case <-tmr:
			return
		default:
			main <- i
		}
	}

}

func reader(c <-chan int) {
	for v := range c {
		fmt.Println("Воркер прочитал :", v)
	}
}

func timer(tmr chan<- int, sec int) {
	time.Sleep(time.Second * time.Duration(sec))
	close(tmr)
}

func Test() {
	main := make(chan int)
	tmr := make(chan int)

	launch(main, tmr)
}
