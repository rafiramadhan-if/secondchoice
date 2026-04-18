package main

import "fmt"

const NMAX int = 127

type tabel [127]rune

func isiArray(t *tabel, n *int) {
	*n = 0
	var ch rune
	fmt.Printf("Teks : ")
	for {
		fmt.Scanf("%c", &ch)
		if ch == '.' || *n >= NMAX {
			break
		}
		if ch != '\n' && ch != ' ' {
			t[*n] = ch
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var balik tabel
	for i := 0; i < n; i++ {
		balik[i] = t[i]
	}
	balikanArray(&balik, n)
	for i := 0; i < n; i++ {
		if t[i] != balik[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	fmt.Print("Reverse teks : ")
	balikanArray(&tab, m)
	cetakArray(tab, m)

	balikanArray(&tab, m)

	fmt.Print("Palindrom     ? ")
	if palindrom(tab, m) {
		fmt.Println("Kata ini adalah palindrom")
	} else {
		fmt.Println("Kata ini bukan palindrom")
	}
}
