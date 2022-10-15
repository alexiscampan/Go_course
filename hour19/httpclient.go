package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	// response, err := http.Get("https://ifconfig.co")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer response.Body.Close()
	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", body)

	// postData := strings.NewReader(`{"some":"json"`)
	// response, err := http.Post("https://httpbin.org/post", "application/json", postData)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer response.Body.Close()
	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", body)

	// client := &http.Client{}
	// request, err := http.NewRequest("GET", "https://ifconfig.co", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// response, err := client.Do(request)
	// defer response.Body.Close()
	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", body)

	debug := os.Getenv("DEBUG")
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://ifconfig.co", nil)
	request.Header.Add("Accept", "application/json")

	if err != nil {
		log.Fatal(err)
	}
	if debug == "1" {
		debugRequest, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", debugRequest)
	}
	response, err := client.Do(request)
	defer response.Body.Close()
	if debug == "1" {
		debugRequest, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", debugRequest)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)

}
