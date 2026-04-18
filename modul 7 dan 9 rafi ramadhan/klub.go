package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	var hasil []string
	pertandingan := 1

	fmt.Println("\nMasukkan skor untuk setiap pertandingan (skor negatif untuk berhenti):")

	for {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			fmt.Println("\nPertandingan selesai.")
			break
		}

		
		if skorA > skorB {
			hasil = append(hasil, klubA)
			fmt.Printf("=> %s menang!\n", klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
			fmt.Printf("=> %s menang!\n", klubB)
		} else {
			hasil = append(hasil, "Draw")
			fmt.Println("=> Hasil imbang (Draw)")
		}
		pertandingan++
	}

	
	fmt.Println("\n--- Rekap Hasil Pertandingan ---")
	for i, h := range hasil {
		fmt.Printf("Hasil %d : %s\n", i+1, h)
	}
}
