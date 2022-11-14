package t10

import (
	"fmt"
	"strconv"
	"strings"
)

type tempMap map[string][]float64

func groupT(arrT []string) tempMap {
	tGrp := make(tempMap)

	for _, v := range arrT {
		t, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println(err)
		}
		if t < 0 {
			tGrp["-"+strconv.Itoa((int(t)/10)*10*-1)] = append(tGrp["-"+strconv.Itoa((int(t)/10)*10*-1)], t)
		} else {
			tGrp[strconv.Itoa((int(t)/10)*10)] = append(tGrp[strconv.Itoa((int(t)/10)*10)], t)
		}

	}
	return tGrp
}

func printTempMap(tGrp tempMap) {
	for k, v := range tGrp {
		fmt.Printf("group: %s, temp: %.1f\n", k, v)
	}
}

func Test() {
	strT := "-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5"
	fmt.Println(strT)
	fmt.Println()

	arrT := strings.Split(strT, ", ")
	mapT := groupT(arrT)

	printTempMap(mapT)

}
