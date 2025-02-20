package tasks

import "fmt"

func Pola2(n int) {
	// bagian atas - tengah
	for i := 0; i < n/2+1; i++ {
		// cetak spasi
		for s := 0; s < i; s++ {
			fmt.Print("\t")
		}
		// cetak angka ganjil sesuai dgn baris
		for j := 0; j < n-2*i; j++ {
			fmt.Print(2*(j+i)+1, "\t")
		}
		fmt.Println()
	}

	// bagian bawah
	for i := n/2 - 1; i >= 0; i-- {
		// cetak spasi
		for s := 0; s < i; s++ {
			fmt.Print("\t")
		}
		// cetak angka ganjil
		for j := 0; j < n-2*i; j++ {
			fmt.Print(2*(j+i)+1, "\t") // cetak ganjil sesuai dgn baris sebelumnya
		}
		fmt.Println()
	}
}
