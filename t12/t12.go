package t12

import "fmt"

type Sundries struct {
	name string
}

func makeLotsOfSundries(str [5]string) []Sundries {
	var lots = make([]Sundries, 0)
	for _, v := range str {
		sund := Sundries{name: v}
		lots = append(lots, sund)
	}
	return lots
}

func Test() {
	str := [5]string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Последовательность строк:", str)
	fmt.Println("Множество:", makeLotsOfSundries(str))
}
