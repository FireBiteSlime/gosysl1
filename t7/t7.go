package t7

import (
	"fmt"
	"strconv"
	"sync"
)

func produce(m map[string]string, i int, mutex *sync.Mutex) {
	mutex.Lock()
	m[strconv.Itoa(i)] = "map with key: " + strconv.Itoa(i)
	mutex.Unlock()
}

func Test() {
	var n int
	fmt.Println("Введите кол-во мап:")
	fmt.Scanf("%d\n", &n)
	val := map[string]string{}
	var mutex sync.Mutex
	for i := 0; i < n; i++ {
		produce(val, i+1, &mutex)
	}

	for _, v := range val {
		mutex.Lock()
		fmt.Println(v)
		mutex.Unlock()
	}

}
