package main

import "fmt"

func soal1(){
	var b1,b2, h float64
	var tanda string

	fmt.Scan(&b1, &tanda, &b2)

	switch {	
	case tanda == "+" :
		h = b1 + b2 
		fmt.Printf("%.1f\n", h)
	case tanda == "-" :
		h = b1 - b2 
		fmt.Printf("%.1f\n", h)
	case tanda == "/" :
		h = b1 / b2 
		fmt.Printf("%.3f\n", h)
	case tanda == "*" :
		h = b1 * b2 
		fmt.Print(h)
	}
}

func soal2(){
	var l int 
	fmt.Scan(&l)

	switch {
	case l >= 1 && l <= 10 :
		fmt.Print("Pemula")
	case l >= 11 && l <= 20 :
		fmt.Print("Menengah")
	case l >= 21 && l <= 30 :
		fmt.Print("Ahli")
	case l > 30 :
		fmt.Print("Master")
	default :
		fmt.Print("Level tidak invalid")
	}
}

func soal3(){
	var bil int
	fmt.Scan(&bil)
	
	switch  {
	case bil < 10 :
		fmt.Print("Satuan")
	case bil < 100 :
		fmt.Print("Puluhan")
	case bil < 1000 :
		fmt.Print("Ratusan")
	case bil < 10000 :
		fmt.Print("Ribuan")
	case bil < 100000 :
		fmt.Print("Puluhan ribu")
	case bil < 1000000 :
		fmt.Print("Ratusan ribu")
	}

}


func main() {
	soal1()
}

