package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type Movie struct {
	Name   string
	Rating float64
}

type Sphere struct {
	Radius float64
}

type Triangle struct {
	base   float64
	height float64
}

type Robot interface {
	PowerOn() error
	// Talk() error
}

type Taxi interface {
	IsAccessible() error
	CountPassangers() error
}

type T850 struct {
	Name string
}

type R2D2 struct {
	Broken bool
}

func (m *Movie) summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + ", " + r
}

func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}

func (s *Sphere) Volume() float64 {
	radiusCubed := s.Radius * s.Radius * s.Radius
	return (float64(4) / float64(3)) * math.Pi * radiusCubed
}

func (t *Triangle) area() float64 {
	return 0.5 * (t.base * t.height)
}

func (t Triangle) changeBase(f float64) {
	t.base = f
	return
}

func (a *T850) PowerOn() error {
	return nil
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}

// func (a *T850) Talk() error {
// 	return errors.New("Hello my friend")
// }

func Boot(r Robot) error {
	return r.PowerOn()
}

func main() {
	m := Movie{
		Name:   "Spiderman",
		Rating: 3.2,
	}
	fmt.Println(m.summary())

	s := Sphere{
		Radius: 5,
	}
	fmt.Println(s.SurfaceArea())
	fmt.Println(s.Volume())

	// t := Triangle{base: 3, height: 1}
	// fmt.Println(t.area())
	// t.changeBase(4)
	// fmt.Println(t.base)

	t := T850{
		Name: "The Terminator",
	}

	r := R2D2{
		Broken: true,
	}

	err := Boot(&r)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}

	err = Boot(&t)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}

}
