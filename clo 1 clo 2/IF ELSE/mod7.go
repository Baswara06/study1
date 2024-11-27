package main
import"fmt"

func soal1(){
	var suhu int
	var temp string
	
	fmt.Scan(&suhu, &temp)
	
	if temp == "Celcius" {
		fmt.Printf ("Suhu dalam Farenheit adalah %d Farenheit\n", ((suhu * 9)/5)+32 )
	} else if temp == "Farenheit"{
		fmt.Printf("Suhu dalam celcius adalah %d Celcius\n", (suhu-32) * 5/9)
	} else { 
		fmt.Println("masukan input dengan benar yang berupa suhu Farenheit dan Celcius")
	}
}

func soal2(){
	var kartu, SuratIzin bool
	fmt.Scan(&kartu, &SuratIzin)

	if kartu || SuratIzin {
		fmt.Println("Diizinkan masuk")
	} else {
		fmt.Println("Tidak boleh masuk")
	}
}

func soal3(){
	var es, jumlah, karung int
	fmt.Scan(&es, &jumlah)
	
	karung = (es * jumlah + 999)/1000
	fmt.Printf("%d karung", karung)
}

func main (){
	soal1()
}