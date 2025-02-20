package tasks

import "fmt"

func Task8(n int) {
	start := 1 // nilai awal angka
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("\t")
		}
		// nentuin jumlah angka yg harus diprint di baris ini
		count := n - i // supaya sgeitiga kebalik

		// cek baris genap/ganjil buat nentuin turun atau naik
		if i%2 == 0 {
			// baris genap: angka naik (ganjil)
			for j := 0; j < count; j++ {
				fmt.Print(start, "\t ")
				start += 2
			}
		} else {
			// baris ganjil: angka turun (ganjil)
			down := start + (count-1)*2 // nentuin angka pertama buat bikin menurun
			for j := 0; j < count; j++ {
				fmt.Print(down, "\t ")
				down -= 2
			}
			start += count * 2 // update start untuk baris berikutnya
		}
		fmt.Println()
	}
}
