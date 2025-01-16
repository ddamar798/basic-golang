package main
import "fmt"

func main()  {
	// for i:=0; i < 6; i++ {
	// 	fmt.Println("Angka",i)
	// }


	 var i = 0
	// for i < 10 {
	// 	fmt.Println("Angka",i)
	// 	i++
	// }


	for {
		fmt.Println("angka",i)
		i++
		if i == 10 {
			break
		}
	}
}