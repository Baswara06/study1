package main

//program menghitung untung penjual
import "fmt"

func main() {
	//kamus
	var modal1, modal2, modal3 int

	//algoritma
	fmt.Scanln(&modal1, &modal2, &modal3)
	untung1 = int(float64(5)/100*float64(modal1) + float64(modal1))
	untung2 = int(float64(5)/100*float64(modal2) + float64(modal2))
	untung3 = int(float64(5)/100*float64(modal3) + float64(modal3))

	fmt.Println(untung1, untung2, untung3)

}
