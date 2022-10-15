package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// fmt.Println(time.Now())
	// time.Sleep(3 * time.Second)
	// fmt.Println("I am awake")
	// fmt.Println("You have two seconds to calculate 19*4")
	// for {
	// 	select {
	// 	case <-time.After(2 * time.Second):
	// 		fmt.Println("Time's up! The answer is 74. Did you get it?")
	// 		return
	// 	}
	// }

	// c := time.Tick(5 * time.Second)
	// for t := range c {
	// 	fmt.Printf("The time is now %v\n", t)
	// }

	// s := "2006-02-02T15:04:05+07:00"
	// t, err := time.Parse(time.RFC3339, s)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(t)

	// fmt.Printf("The hour is %v\n", t.Hour())

	s1 := "2017-09-03T18:00:00+00:00"
	s2 := "2017-09-04T18:00:00+00:00"
	today, err := time.Parse(time.RFC3339, s1)
	if err != nil {
		log.Fatal(err)
	}
	tomorrow, err := time.Parse(time.RFC3339, s2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(today.After(tomorrow))
	fmt.Println(today.Before(tomorrow))
	fmt.Println(today.Equal(tomorrow))

}
