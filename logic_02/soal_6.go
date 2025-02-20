package logic_02

import go_print_slice "github.com/zaadinda/go-print-slice"

func Soal_6(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)
	start := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			if i%2 == 0 { //genap ditambah ganjil
				result[i][j] = start
				start += 3
			} else {
				result[i][n-1-j] = start // ganjil ditambah genap
				start += 2
			}
		}
	}
	return result
}
