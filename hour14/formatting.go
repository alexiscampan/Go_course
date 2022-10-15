package formatting

import (
	"errors"
	"fmt"
	"math"
)

const Foo string = "constant string"

//Animal specifies an animal
type Animal struct {
	Name string // Name holds the name of a thing
	Age  int    //Age holds the name of a thing
}

// ErrNotAnAnimal is returned if the name field of the Animal struct is Human.
var ErrNotAnAnimal = errors.New("Name is not an animal")

func (a Animal) Hello() (string, error) {
	if a.Name == "Human" {
		return "", ErrNotAnAnimal
	}
	s := "Hello " + a.Name
	return s, nil
}

func main() {
	var f float64 = 9
	fmt.Println(math.Sqrt(f))
	fmt.Println(Foo)
	a_string := "hello"
	fmt.Println(a_string)
}
