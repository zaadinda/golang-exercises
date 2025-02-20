package main

import (
	"fmt"
	go_print_slice "github.com/zaadinda/go-print-slice"
	"github.com/zaadinda/golang-exercises/logic_02"
	"github.com/zaadinda/golang-exercises/tasks"
)

func main() {
	//Logic 1 - No.1 - task1
	fmt.Println("Logic 1 - N0. 1")
	tasks.Task1()
	fmt.Println()

	//Logic 1 - No. 6 - task2
	fmt.Println("Logic 1 - N0. 6")
	tasks.KurangTiga(10)
	fmt.Println()

	// Logic 1 - No. 8 - task6
	fmt.Println("Logic 1 - NO. 8")
	tasks.Task6(10)
	fmt.Println()
	tasks.Task6(11)
	fmt.Println()

	// Logic 1 - No. 9
	fmt.Println("Logic 1 - NO. 9")
	tasks.Task7(10)
	fmt.Println()
	tasks.Task7(11)
	fmt.Println()

	//Logic 1 - No. 11 - task5
	fmt.Println("Logic 1 - NO. 11")
	tasks.Nomer11()
	fmt.Println()

	//Logic 2 - No. 1 - task3
	fmt.Println("Logic 2 - N0. 1")
	tasks.Pola1(9)
	fmt.Println()

	// Logic 2 - No. 6 - soal6
	fmt.Println("Logic 2 - N0. 6")
	slice2_6 := logic_02.Soal_6(9)
	go_print_slice.Print2D(slice2_6)
	fmt.Println()

	// Logic 2 - No. 13 - task4
	fmt.Println("Logic 2 - N0. 13")
	tasks.Pola2(9)
	fmt.Println()

	// Logic 2 - NO.7 - task9
	fmt.Println("Logic 2 - NO. 7")
	tasks.Task9(9)
	fmt.Println()

	// Logic 3 - NO. 2 - task8
	fmt.Println("Logic 3 - N0. 2")
	tasks.Task8(9)
	fmt.Println()

}
