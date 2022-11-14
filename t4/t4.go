package t4

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func launch(main chan string, cls chan os.Signal) {
	var amount int

	fmt.Println("Введите колличество воркеров: ")
	fmt.Scanf("%d\n", &amount)

	for i := 0; i < amount; i++ {

		go worker(main, i)
	}

	for {
		str := Randstr(3)
		select {
		case main <- str:
		case <-cls:
			time.Sleep(1 * time.Second)
			close(main)
			fmt.Println("Конец работы программы")
			return
		}
	}

}

func worker(c <-chan string, i int) {
	for v := range c {
		time.Sleep(1 * time.Second)
		fmt.Println("Воркер номер: ", i+1, " Прочитал: ", v)
	}
}

func Randstr(lg int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	str := make([]rune, lg)
	for i := range str {
		str[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(str)
}

func Test() {
	rand.Seed(time.Now().UnixNano()) // рандомизируем генератор случайных чисел

	cls := make(chan os.Signal, 1)
	signal.Notify(cls, syscall.SIGINT) // отлов сигнала завершения программы(CTRL + C)

	main := make(chan string)
	launch(main, cls)
}
