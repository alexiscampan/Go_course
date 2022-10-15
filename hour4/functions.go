package main

import (
	"fmt"
)

func getPrize() (int, string) {
	i := 2
	s := "goldfish"
	return i, s
}

func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func sayHi() (x, y string) {
	x = "Hello"
	y = "World"
	return
}

// recursivity
func feedMe(portion int, eaten int) int {
	eaten = portion + eaten
	if eaten >= 5 {
		fmt.Printf("I'm full! I've eaten %d\n", eaten)
		return eaten
	}
	fmt.Printf("I'm still hungry! I've eaten %d\n", eaten)
	return feedMe(portion, eaten)
}

// function as argument

func anotherFunction(f func() string) string {
	return f()
}

// function take 2 return 3

func return3func(x int, y int) (int, int, int) {
	z := x + y
	return x, y, z
}

func main() {
	quantity, price := getPrize()
	fmt.Printf("You won %v %v\n", quantity, price)

	result := sumNumbers(1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("The result is %v\n", result)

	fmt.Println(sayHi())

	fmt.Println(feedMe(1, 0))

	fn := func() string {
		return "function called"
	}
	fmt.Println(anotherFunction(fn))

	fmt.Println(return3func(2, 6))
}
