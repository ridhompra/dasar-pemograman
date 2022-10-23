package main

import "fmt"

func main() {

	var tabungan, n, hasil int

	fmt.Scan(&tabungan, &n)
	for i := 0; i <= n-1; i++ {
		hasil += tabungan + 2500*i
	}
	fmt.Println(hasil)

}
