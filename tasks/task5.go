package tasks

import "fmt"

func Nomer11() {
	n := 10
	num := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print("buzz ")
		} else {
			fmt.Print(num, " ")
			if 1 == i {
				num *= 3
			} else {
				num *= 2
			}
		}
	}

}
