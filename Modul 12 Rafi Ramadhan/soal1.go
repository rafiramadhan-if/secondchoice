package main
 
import "fmt"
 
const MAXCALON = 20
 
func main() {
	var suara [MAXCALON + 1]int
	var masuk, sah, pilih int
 
	fmt.Scan(&pilih)
	for pilih != 0 {
		masuk++
		if pilih >= 1 && pilih <= MAXCALON {
			sah++
			suara[pilih]++
		}
		fmt.Scan(&pilih)
	}
 
	fmt.Printf("Suara masuk: %d\n", masuk)
	fmt.Printf("Suara sah: %d\n", sah)
 
	var i int = 1
	for i <= MAXCALON {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
		i++
	}
}
