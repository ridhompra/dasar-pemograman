package main

import "fmt"

func main() {
	var n, kakak, adik, hasil int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&adik, &kakak)
		hasil = adik - kakak
		if hasil < 0 {
			hasil = (-hasil)
		}
		if adik == kakak*2 || adik < kakak && kakak%2 == 1 || hasil == 1 {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	}
}
