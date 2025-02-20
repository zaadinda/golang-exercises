package tasks

import "fmt"

func Task6(n int) {
	mid := n / 2 // nilai tengah dari n

	for i := 1; i <= n; i++ {
		if i <= mid { // jika i ada di tengah, print angka genap naik i * 2
			fmt.Print(i*2, "\t")
		} else { // kalo udah lewat tengah, print angke genap turun
			fmt.Print((n-i+1)*2, "\t")
		}
	}
}
