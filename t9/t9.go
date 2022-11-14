package t9

import "sync"

func readV(arr [5]int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	for _, v := range arr {
		ch <- v
	}
}

func multiplV(arr <-chan int, arrSq chan<- int, wg *sync.WaitGroup) {
	for {
		val, err := <-arr
		if !err {
			close(arrSq)
			wg.Done()
			return
		}
		arrSq <- (val * 2)
	}

}

func printV(ch <-chan int, wg *sync.WaitGroup) {
	for {
		val, err := <-ch
		if !err {
			wg.Done()
			return
		}
		println(val)
	}
}

func Test() {

	arr := [5]int{1, 2, 3, 4, 5}
	arrChan := make(chan int)
	arrPowChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	go readV(arr, arrChan, &wg)
	go multiplV(arrChan, arrPowChan, &wg)
	go printV(arrPowChan, &wg)

	wg.Wait()
}
