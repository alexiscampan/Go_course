package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper() finished")
}

func responseTime(url string) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	elapsed := time.Since(start).Seconds()
	fmt.Printf("%s took %v seconds \n", url, elapsed)

}

func main() {
	// go slowFunc()
	// fmt.Println("I am not shown until slowFunc completes")
	// time.Sleep(time.Second * 3)

	urls := make([]string, 3)
	urls[0] = "https://www.usa.gov/"
	urls[1] = "https://www.gov.uk/"
	urls[2] = "http://www.gouvernement.fr/"
	for _, u := range urls {
		go responseTime(u)
	}
	time.Sleep(time.Second * 5)
}
