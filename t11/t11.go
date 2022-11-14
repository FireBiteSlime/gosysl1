package t11

import (
	"fmt"
	"math/rand"
	"time"
)

func printCoinciden(arrF []int, arrS []int) {
	for _, fv := range arrF {
		for _, sv := range arrS {
			if fv == sv {
				println(fv)
			}
		}
	}
}

func Test() {
	rand.Seed(time.Now().UnixNano())
	arrF := make([]int, 0)
	arrS := make([]int, 0)
	for i := 0; i < rand.Intn(15+1); i++ {
		arrF = append(arrF, rand.Intn(5)+1)
	}
	for i := 0; i < rand.Intn(15+1); i++ {
		arrS = append(arrS, rand.Intn(5)+1)
	}
	fmt.Println("Первое неупорядоченное множество: ", arrF)
	fmt.Println()
	fmt.Println("Второе неупорядоченное множество: ", arrS)
	fmt.Println()

	printCoinciden(arrF, arrS)

}
