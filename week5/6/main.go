package main

import "fmt"

func main() {
	var bilangan int
	for i := 2; i <= 50; i++ {
		bilangan += i
	}
	fmt.Println(bilangan)
}
