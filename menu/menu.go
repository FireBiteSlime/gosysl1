package menu

import (
	"fmt"
	"l1/t1"
	"l1/t10"
	"l1/t11"
	"l1/t12"
	"l1/t13"
	"l1/t14"
	"l1/t15"
	"l1/t16"
	"l1/t17"
	"l1/t18"
	"l1/t19"
	"l1/t2"
	"l1/t20"
	"l1/t21"
	"l1/t22"
	"l1/t23"
	"l1/t24"
	"l1/t25"
	"l1/t26"
	"l1/t3"
	"l1/t4"
	"l1/t5"
	"l1/t6"
	"l1/t7"
	"l1/t8"
	"l1/t9"
)

func Show() {
	var item int
	for {
		//time.Sleep(1 * time.Second)
		fmt.Println("Выберите номер задачи от 1 до 26")
		fmt.Scanf("%d\n", &item)
		//ClearScreen()
		Order(item)
		fmt.Println()
	}
}

func Order(taskid int) {
	switch taskid {
	case 1:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 1")
		fmt.Println()
		t1.Test()
		fmt.Println("------------------------------------------")
	case 2:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 2")
		fmt.Println()
		t2.Test()
		fmt.Println("------------------------------------------")
	case 3:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 3")
		fmt.Println()
		t3.Test()
		fmt.Println("------------------------------------------")
	case 4:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 4")
		fmt.Println()
		t4.Test()
		fmt.Println("------------------------------------------")
	case 5:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 5")
		fmt.Println()
		t5.Test()
		fmt.Println("------------------------------------------")
	case 6:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 6")
		t6.Test()
		fmt.Println("------------------------------------------")
	case 7:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 7")
		fmt.Println()
		t7.Test()
		fmt.Println("------------------------------------------")
	case 8:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 8")
		t8.Test()
		fmt.Println("------------------------------------------")
	case 9:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 9")
		t9.Test()
		fmt.Println("------------------------------------------")
	case 10:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 10")
		fmt.Println()
		t10.Test()
		fmt.Println("------------------------------------------")
	case 11:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 11")
		fmt.Println()
		t11.Test()
		fmt.Println("------------------------------------------")
	case 12:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 12")
		fmt.Println()
		t12.Test()
		fmt.Println("------------------------------------------")
	case 13:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 13")
		fmt.Println()
		t13.Test()
		fmt.Println("------------------------------------------")
	case 14:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 14")
		fmt.Println()
		t14.Test()
		fmt.Println("------------------------------------------")
	case 15:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 15")
		fmt.Println()
		t15.Test()
		fmt.Println("------------------------------------------")
	case 16:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 16")
		fmt.Println()
		t16.Test()
		fmt.Println("------------------------------------------")
	case 17:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 17")
		fmt.Println()
		t17.Test()
		fmt.Println("------------------------------------------")
	case 18:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 18")
		fmt.Println()
		t18.Test()
		fmt.Println("------------------------------------------")
	case 19:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 19")
		fmt.Println()
		t19.Test()
		fmt.Println("------------------------------------------")
	case 20:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 20")
		fmt.Println()
		t20.Test()
		fmt.Println("------------------------------------------")
	case 21:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 21")
		fmt.Println()
		t21.Test()
		fmt.Println("------------------------------------------")
	case 22:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 22")
		fmt.Println()
		t22.Test()
		fmt.Println("------------------------------------------")
	case 23:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 23")
		fmt.Println()
		t23.Test()
		fmt.Println("------------------------------------------")
	case 24:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 24")
		fmt.Println()
		t24.Test()
		fmt.Println("------------------------------------------")
	case 25:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 25")
		fmt.Println()
		t25.Test()
		fmt.Println("------------------------------------------")
	case 26:
		fmt.Println("------------------------------------------")
		fmt.Println("Task 26")
		fmt.Println()
		t26.Test()
		fmt.Println("------------------------------------------")
	default:
		return
	}

}

func ClearScreen() {
	fmt.Print("\x1b[40m")
	fmt.Print("\x1b[2J")
}
