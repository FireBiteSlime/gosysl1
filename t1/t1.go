package t1

import "fmt"

type Human struct {
	age  int
	male string
}

func (h *Human) AddHuman(age int, male string) {
	h.age = age
	h.male = male
}

type Action struct {
	Human
}

type Action1 struct {
	h Human
}

func Test() {
	var zero Action
	var one Action1

	fmt.Println("Безымянное встраивание в структуру Action:")
	zero.AddHuman(10, "men")
	fmt.Println(zero)
	fmt.Println("Встраивание с именем в структуру Action1:")
	one.h.AddHuman(14, "women")
	fmt.Println(one)

}
