package t25

import (
	"fmt"
	"time"
)

func mySleep(t int) {
	<-time.After(time.Second * time.Duration(t))
}

func Test() {
	fmt.Println("Введите время сна:")
	var t int
	fmt.Scanf("%d\n", &t)
	mySleep(t)
}
