package tasks

import "fmt"

func KurangTiga(n int) {
	for i := n * 3; i >= 3; i -= 3 {
		fmt.Print(i, "\t")
	}
	fmt.Println()
}
