package tasks

import "fmt"

func Task9(n int) {
	start := 1

	for i := 0; i < n; i++ {
		// print spasi ngebentuk segitiga terbalik
		for j := 0; j < i; j++ {
			fmt.Print("\t ")
		}
		// cetak angka
		fmt.Println(start)

		start += 2
	}
}
