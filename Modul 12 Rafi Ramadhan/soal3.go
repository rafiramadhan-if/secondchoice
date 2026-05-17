package main
 
import "fmt"
 
const NMAX = 1000000
 
var data [NMAX]int
 
func isiArray(n int) {
	var i int = 0
	for i < n {
		fmt.Scan(&data[i])
		i++
	}
}
 
func posisi(n, k int) int {
	var found int = -1
	var kr int = 0
	var kn int = n - 1
	var med int
 
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if k < data[med] {
			kn = med - 1
		} else if k > data[med] {
			kr = med + 1
		} else {
			found = med
		}
	}
	return found
}
 
func main() {
	var n, k int
	fmt.Scan(&n, &k)
	isiArray(n)
 
	hasil := posisi(n, k)
	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}
