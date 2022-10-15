package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	// b := make([]byte, 0)
	// err := ioutil.WriteFile("example02.txt", b, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	for _, file := range files {
		fmt.Println(file.Mode(), file.Name())
	}
}
