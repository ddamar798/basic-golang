package main
import "fmt"

func main() {
	var buah[4]string
	buah[0] = "Apel"
	buah[1] = "Jeruk"
	buah[2] = "Mangga"
	buah[3] = "Anggur"
	fmt.Println("saya memiliki ", len(buah), " jenis buah")
	fmt.Println("saya memiliki buah ", buah[2])
}