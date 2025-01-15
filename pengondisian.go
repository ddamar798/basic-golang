package main
import "fmt"

func main(){
	var nilai int

	fmt.Println("Masukan Nilai : ")
	fmt.Scanln(&nilai)

	if nilai <10 && nilai >=7{
		fmt.Println("Sempurna!")
	} else if nilai <7 && nilai >=5 {
		fmt.Println("Cukup!")
	} else if nilai <5 {
		fmt.Println("Kurang!")
	}
}