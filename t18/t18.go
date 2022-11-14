package t18

import (
	"fmt"
	"sync"
)

type Counter struct {
	num int
	sync.Mutex
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()

	c.num += 1
}

func (c *Counter) Value() int {
	return c.num
}

func Do(cnt *Counter, finish chan struct{}) {
	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(num int, cnt *Counter, wg *sync.WaitGroup) {
			defer wg.Done()

			fmt.Printf("Worker %d starting\n", num)
			cnt.Inc()
			fmt.Printf("Worker %d done\n", num)
		}(i, cnt, &wg)
	}

	wg.Wait()
	close(finish)
}

func Test() {
	cnt := &Counter{
		num: 0,
	}

	finish := make(chan struct{})

	go Do(cnt, finish)

	select {
	case <-finish:
		fmt.Println("Счетчик:", cnt.Value())
	}
}
