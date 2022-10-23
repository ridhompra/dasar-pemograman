package main

import "fmt"

func main() {
	var bilangan int
	for i := 1; i <= 1000; i++ {
		bilangan += i
	}
	fmt.Println(bilangan)
}
