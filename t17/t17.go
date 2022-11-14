package t17

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func Test() {
	rand.Seed(time.Now().Unix())
	vals := make([]int, 0)
	for i := 0; i < rand.Intn(10)+13; i++ {
		vals = append(vals, rand.Intn(10)+1)
	}
	sort.Ints(vals)
	var x int
	fmt.Println("Массив:", vals)
	sort.Ints(vals)
	fmt.Println("Введите элемент:")
	fmt.Scanf("%d\n", &x)
	i := sort.Search(len(vals), func(i int) bool { return vals[i] >= x })
	if i < len(vals) && vals[i] == x {
		fmt.Printf("Найден: %d под номером: %d\n", x, i)
	} else {
		fmt.Println("Элемент не найден")
	}
}
