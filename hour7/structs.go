package main

import "fmt"

type Movie struct {
	Name   string
	Rating float32
}

type SuperHero struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Number int
	Street string
	City   string
}

type Alarm struct {
	Time  string
	Sound string
}

func NewAlarm(time string) Alarm {
	a := Alarm{
		Time:  time,
		Sound: "Klaxon",
	}
	return a
}

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	m := Movie{
		Name:   "Salut mec",
		Rating: 10,
	}
	fmt.Printf("%+v\n", m)

	e := SuperHero{
		Name: "frero",
		Age:  20,
		Address: Address{
			Number: 1007,
			Street: "rue du yes",
			City:   "montpellier",
		},
	}

	fmt.Printf("%+v\n", e)

	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	// b := Drink{
	// 	Name: "Lemonade",
	// 	Ice:  true,
	// }
	// if a == b {
	// 	fmt.Println("same")
	// }
	b := &a
	b.Ice = false
	fmt.Printf("%+v\n", *b)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &a)
}
