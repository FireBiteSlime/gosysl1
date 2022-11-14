package t23

import "fmt"

func Test() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := make([]int, len(a))
	copy(b, a)
	fmt.Println(b)
	fmt.Println("Введите номер элемента для удаления:")
	n := 0
	fmt.Scanf("%d\n", &n)
	if n < len(a) {
		a[n] = a[len(a)-1]
		a[len(a)-1] = 0
		a = a[:len(a)-1]
		fmt.Println("Удаление с изменением порядка: ", a)

		copy(b[n:], b[n+1:])
		b[len(b)-1] = 0
		b = b[:len(b)-1]
		fmt.Println("Удаление без изменения порядка:", b)
	}
}
