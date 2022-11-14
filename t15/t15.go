package t15

import (
	"fmt"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func createHugeString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//var justString string //глобальная переменная делает код сложным для тестирования(+ может помешать встроенной конкуретности)

func someFunc(justString *string) {
	//v := createHugeString(1 << 10) // не понятно что происходит внутри функции и что она вызывает, так же использование оператора << в данном случае вредит удобочитаемости когда и резервирует больше памяти и мощности машины
	v := createHugeString(100)
	//if len(v) >= 100 {             // нет смысла проверять длинну, если функция createHugeString вернёт строку определённой длинны
	//*justString = v[:100] //мы не знаем что нам вернула createHugeString может быть ошибка,так же исходя из логики мы запросили переменную длинной 1024 символа, а используем только 100
	*justString = v
	//}

	fmt.Println("Result:", *justString)
}

func Test() {
	justString := ""
	rand.Seed(time.Now().UnixNano())
	someFunc(&justString)
}
