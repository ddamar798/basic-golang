package main
import "fmt"

func main(){
	var jumlah int

	fmt.Println("Isikan Jumlah barang :  ")
	fmt.Scanln(&jumlah) // syntax "scan" digunakan untuk membuat input.

	switch jumlah {
	case 10:
		fmt.Println("Jumlah Cukup")
	case 6 :
	fmt.Println("Jumlah Sedikit, Lakukan Isi Ulang!")
	}
}