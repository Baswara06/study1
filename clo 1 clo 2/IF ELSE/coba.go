package main

import "fmt"

func main(){
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
