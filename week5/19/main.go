package main

import (
	"fmt"
	"math"
)

func main() {
	var n, data int
	var hasil, harmonic float64
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data)
		harmonic += float64(1) / float64(data)
	}
	hasil = (float64(n) / harmonic)
	fmt.Println(roundFloat(hasil, 2))

}
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
