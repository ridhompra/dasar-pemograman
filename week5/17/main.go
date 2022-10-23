package main

import "fmt"

func main() {
	var n, goal, total int
	var bool bool
	fmt.Scan(&n, &goal)

	for i := 0; i < n; i++ {
		total += goal
	}
	if total >= 10 {
		bool = true
	} else {
		bool = false
	}
	fmt.Println(bool)
}
