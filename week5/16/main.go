package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	var hasil float64
	fmt.Scan(&n, &m)
	for i := n; i <= m; i++ {
		floati := float64(i)
		hasil += 4 / floati
	}
	fmt.Println(roundFloat(hasil, 2))
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
