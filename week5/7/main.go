package main

import "fmt"

func main() {
	var n, shopbyday, hasil int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&shopbyday)
		hasil += shopbyday
	}
	fmt.Println(hasil)
}
