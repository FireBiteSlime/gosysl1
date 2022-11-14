package t26

import (
	"fmt"
	"strings"
)

func Test() {
	str := ""
	fmt.Println("Введите слово:")
	fmt.Scanf("%s\n", &str)
	fmt.Printf("%s - ", str)
	str = strings.ToLower(str)
	uniStr := []rune(str)
	f := true
	for i := 0; i < len(uniStr)-1; i++ {
		if uniStr[i] == uniStr[i+1] {
			f = false
		}
	}

	fmt.Printf("%t\n", f)
}
