package main

import "fmt"

func main() {
	var n, data, sum, jam, menit int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data)
		sum += data
		jam = sum / 60
		menit = sum - (jam * 60)
	}
	fmt.Println(jam, menit)
}
