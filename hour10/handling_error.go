package main

import (
	"fmt"
)

func Half(numberToHalf int) (int, error) {
	if numberToHalf%2 != 0 {
		return -1, fmt.Errorf("Cannot half %v", numberToHalf)
	}
	return numberToHalf / 2, nil
}

func main() {
	// file, err := ioutil.ReadFile("foo.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("%s", file)

	// err := errors.New("Something went wrong")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// name, role := "Richard Jupp", "Drummer"
	// err := fmt.Errorf("The %v %v quit", role, name)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// n, err := Half(19)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(n)

	fmt.Println("This is executed")
	panic("Oh no. I can do more. Goodbye.")
	fmt.Println("This is not executed")
}
