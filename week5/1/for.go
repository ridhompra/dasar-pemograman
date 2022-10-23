package main

import "fmt"

func main() {
	var bilangan int
	for i := 1; i <= 100; i++ {
		bilangan += i
	}
	fmt.Println(bilangan)
}
