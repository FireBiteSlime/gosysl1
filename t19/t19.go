package t19

import (
	"fmt"
)

func Test() {
	var str string
	fmt.Println("Введите слово:")
	fmt.Scanf("%s\n", &str)
	uniStr := []rune(str)
	var res string
	for i := len(uniStr) - 1; i >= 0; i-- {
		res = res + string(uniStr[i])
	}
	fmt.Println(res)
}
