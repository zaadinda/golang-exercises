package tasks

import "fmt"

func Pola1(n int) {
	for i := 0; i < n; i++ { // ngulang sebanyak n baris
		for j := 0; j < n; j++ { // buat ngisi tiap baris pake angka
			fmt.Print(2*j+1, "\t") // rumus angka ganjil
		}
		fmt.Println() // cetak ke baris baru
	}
}
