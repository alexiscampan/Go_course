package main

import (
	"fmt"
)

func evaluateInt(x int) {
	if x <= 200 {
		fmt.Printf("less than 200")
	} else {
		fmt.Printf("more than 200")
	}
}

func main() {
	// b := true
	// if b {
	// 	fmt.Println("b is true!")
	// }

	// i := 2
	// switch i {
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// case 4:
	// 	fmt.Println("Four")
	// default:
	// 	fmt.Println("error")
	// }

	defer evaluateInt(88)
	fmt.Println("It runs")
}
