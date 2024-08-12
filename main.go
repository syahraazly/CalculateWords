package main

import (
	"fmt"
)

func main() {
	// Kalimat yang akan dipecah
	sentence := "golang is awesome"

	// Membuat map untuk menghitung kemunculan setiap karakter
	charCount := make(map[rune]int)

	// Looping melalui setiap karakter dalam kalimat
	for _, char := range sentence {
		// Mencetak setiap karakter pada baris baru
		fmt.Printf("%c\n", char)

		// Menambah jumlah kemunculan karakter pada map
		charCount[char]++
	}

	// Menampilkan jumlah kemunculan setiap karakter
	fmt.Println("\nJumlah kemunculan setiap karakter:")
	for char, count := range charCount {
		fmt.Printf("'%c' : %d\t", char, count)
	}
}
