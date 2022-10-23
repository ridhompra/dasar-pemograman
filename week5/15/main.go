package main

import "fmt"

func main() {
	var n int
	var message string
	fmt.Scan(&n, &message)
	for i := 0; i < n; i++ {
		fmt.Println(message)
	}
}
