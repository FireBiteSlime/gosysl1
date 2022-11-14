package t20

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Test() {
	fmt.Println("Введите строку:")
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	s := myscanner.Text()
	vals := strings.Fields(s)
	s = ""
	for i := len(vals) - 1; i >= 0; i-- {
		s = s + vals[i] + " "
	}
	fmt.Println(s)
}
