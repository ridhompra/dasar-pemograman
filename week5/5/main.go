package main

import "fmt"

func main() {
	var n, MasukJam, MasukMenit, MasukDetik, KeluarJam, KeluarMenit, KeluarDetik, SumPerDay, SumWorkHour int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&MasukJam, &MasukMenit, &MasukDetik, &KeluarJam, &KeluarMenit, &KeluarDetik)
		SumPerDay = (KeluarJam*3600 + KeluarMenit*60 + KeluarDetik) - (MasukJam*3600 + MasukMenit*60 + MasukDetik)
		// fmt.Println(KeluarJam, KeluarMenit, KeluarDetik)
		// fmt.Println(MasukJam, MasukMenit, MasukDetik)
		// fmt.Println(SumPerDay / 3600)
		SumWorkHour += SumPerDay
		// fmt.Println(SumWorkHour)
	}
	hasil := SumWorkHour / 3600
	// fmt.Println(hasil)
	// fmt.Println(SumWorkHour)
	if hasil >= 40 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
