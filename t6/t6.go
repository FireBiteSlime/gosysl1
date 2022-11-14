package t6

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func waysMenu() {
	var exid int
	var seconds int

	fmt.Println("Введите колличество секунд работы воркеров: ")
	fmt.Scanf("%d\n", &seconds)

	fmt.Println("Введите номер примера:")
	fmt.Println("1 - Завершение с использованием доп.канала")
	fmt.Println("2 - Завершение с использованием указателя")
	fmt.Println("3 - Завершение с использованием sync.WaitGroup")
	fmt.Println("4 - Завершение с использованием Context")
	fmt.Scanf("%d\n", &exid)

	switch exid {
	case 1:
		clsChnl(seconds)
	case 2:
		clsPointr(seconds)
	case 3:
		clsWg(seconds)
	case 4:
		clsCtx(seconds)
	}
}

// //////////////закрытие с помощью доп. канала tmr///////////////
func reader(c <-chan int) {
	for v := range c {
		fmt.Println("Воркер прочитал :", v)
	}
}

func chnTimer(tmr chan<- int, sec int) {
	time.Sleep(time.Second * time.Duration(sec))
	close(tmr)
}

func clsChnl(seconds int) {
	main := make(chan int)
	tmr := make(chan int)

	go chnTimer(tmr, seconds)
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

// ////////////////закрытие с помощью указателя///////////////
func clsPointr(seconds int) {
	stop := true
	go stopByPointr(&stop)
	pointrTimer(&stop, seconds)
	for i := 0; ; i++ {
		if stop {
			fmt.Println(i)
		} else {
			return
		}
	}
}

func stopByPointr(stop *bool) {
	for i := 0; ; i++ {
		if *stop {
			fmt.Println(i)
		} else {
			return
		}
	}
}

func pointrTimer(stop *bool, sec int) {
	time.Sleep(time.Second * time.Duration(sec))
	*stop = false
}

// ////////////////закрытие с помощью sync.WaitGroup///////////////
func clsWg(seconds int) {
	chn := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go stopByWg(chn, &wg)
	wgTimer(chn, seconds)
	close(chn)
	wg.Wait()
}

func stopByWg(c chan int, wg *sync.WaitGroup) {
	for {
		foo, err := <-c
		if !err {
			wg.Done()
			return
		}
		println(foo)
	}
}

func wgTimer(c chan int, seconds int) {
	newTimer := time.NewTimer(time.Second * time.Duration(seconds))
	for i := 0; ; i++ {
		select {
		case <-newTimer.C:
			return
		default:
			c <- i
		}
	}
}

// //////////////закрытие с помощью Context////////////////////////
func clsCtx(seconds int) {
	go stopByCtx(seconds)
}

func stopByCtx(sec int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(sec)*time.Second)
	defer cancel()

	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Воркер завершён")
			return
		default:
			fmt.Println(i)
		}
	}
}

func Test() {
	waysMenu()
}
