package main

import "fmt"

func soal1() {
	var x int
	fmt.Scan(&x)

	if x > 1000 {
		fmt.Print(x - (x * 20 / 100))
	} else if 500 <= x && x >= 100 {
		fmt.Print(x - (x * 15 / 100))
	} else {
		fmt.Print(x - (x * 5 / 100))
	}
}

func soal2() {
	var x int
	fmt.Scan(&x)

	if x < 10 {
		fmt.Print("satuan")
	} else if x < 100 {
		fmt.Print("Puluhan")
	} else if x < 1000 {
		fmt.Print("Ratusan")
	} else if x < 100000 {
		fmt.Print("Puluhan ribu")
	} else if x < 1000000 {
		fmt.Print("Ratusan ribu")
	}
}

func soal3() {
	var tb, bb float64
	var bmi float64
	fmt.Scan(&tb, &bb)

	tb = tb / 100
	bmi = bb / (tb * tb)

	if bmi < 18.5 {
		fmt.Printf("Berat badan kurang")
	} else if bmi >= 18.5 && bmi <= 22.9 {
		fmt.Printf("Berat badan normal ")
	} else {
		fmt.Print("Kelebihan berat badan")
	}

}

func main() {
	soal3()
}
