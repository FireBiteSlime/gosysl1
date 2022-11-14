package t22

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

func Test() {
	rand.Seed(time.Now().Unix())
	a := big.NewFloat(math.Pow(2, 20) + rand.Float64()*(1000-100))
	b := big.NewFloat(math.Pow(2, 20) + rand.Float64()*(1000-100))
	fmt.Printf("Первое число: %f; Второе число: %f\n", a, b)
	var n int
	fmt.Printf("Выберете операцию: \n1) умножение \n2) деление \n3) сложение \n4) вычитание \n")
	fmt.Scanf("%d\n", &n)
	switch n {
	case 1:
		fmt.Printf("%f * %f = %f\n", a, b, new(big.Float).Mul(a, b))
		return
	case 2:
		fmt.Printf("%f / %f = %f\n", a, b, new(big.Float).Quo(a, b))
		return
	case 3:
		fmt.Printf("%f + %f = %f\n", a, b, new(big.Float).Add(a, b))
		return
	case 4:
		fmt.Printf("%f - %f = %f\n", a, b, new(big.Float).Sub(a, b))
		return
	default:
		return
	}
}
