package main
 
import "fmt"
 
const MAXCALON = 20
 
func cariPemenang(suara [MAXCALON + 1]int) (int, int) {
	var ketua, wakil int = -1, -1
	var i int = 1
	for i <= MAXCALON {
		if suara[i] > 0 {
			if ketua == -1 || suara[i] > suara[ketua] {
				wakil = ketua
				ketua = i
			} else if wakil == -1 || suara[i] > suara[wakil] {
				wakil = i
			}
		}
		i++
	}
	return ketua, wakil
}
 
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
 
	ketua, wakil := cariPemenang(suara)
	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
