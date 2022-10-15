package main

import "fmt"

// var s = "Hello World"

func showMemoryAddress(x *int) {
	fmt.Println(x)
	return
}

func main() {
	// var (
	// 	s string = "foo"
	// 	i int    = 4
	// )

	// fmt.Println(s)
	// fmt.Println(i)

	// var i int
	// var f float32
	// var b bool
	// var s string
	// fmt.Printf("%v %v %v %q\n", i, f, b, s)
	// var s string
	// if s == "" {
	// 	fmt.Printf("s not assigned")
	// }

	// fmt.Printf("Print 's' variable from outer block %v\n",
	// 	s)
	// b := true
	// if b {
	// 	fmt.Printf("Printing 'b' variable from outer block %v\n", b)
	// 	i := 42
	// 	if b != false {
	// 		fmt.Printf("Printing 'i' variable from outer block %v\n", i)
	// 	}
	// }

	// s := "Hello World"
	// fmt.Println(&s)

	i := 1
	fmt.Println(&i)
	showMemoryAddress(&i)

}
