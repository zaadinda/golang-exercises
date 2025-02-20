package tasks

import "fmt"

func Task1() {
	n := 10                  // jumlah angka yg mau ditampilin
	for i := 0; i < n; i++ { // perulangan dari 0 sampai n-1
		fmt.Print(2*i+1, " ") // rumus ngehasilin angka ganjil
	}
	fmt.Println()
}
