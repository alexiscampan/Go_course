package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// log.Printf("This is a log message")

	// var errFatal = errors.New("we only just started and we are crashing")
	// log.Fatal(errFatal)

	// f, err := os.OpenFile("example.log",
	// 	os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// log.SetOutput(f)

	// for i := 1; i <= 5; i++ {
	// 	log.Printf("Log iteration %d", i)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Guess the name of my pet to win a prize: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, " ", "", -1)
	fmt.Println(text == "John")
	if text == "John" {
		fmt.Println("You won! You win chocolate!")
	} else {
		fmt.Println("You didn't win. Better luck next time")
	}

}
