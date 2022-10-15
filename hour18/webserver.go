package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// switch r.Header.Get("Accept") {
	// case "application/json":
	// 	w.Header().Set("Content-type", "application/json;charset=utf-8")
	// 	w.Write([]byte(`{"Hello":"World"}`))
	// case "application/xml":
	// 	w.Header().Set("Content-type", "application/xml;charset=utf-8")
	// 	w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?><Message>Hello World</Message>`))
	// default:
	// 	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// 	w.Write([]byte("Hello World\n"))
	// }

	switch r.Method {
	case "GET":
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
		}
		w.Write([]byte("Received a GET request\n"))
	case "POST":
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", reqBody)
		w.Write([]byte("Received a POST request\n"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented) + "\n"))

	}

}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8000", nil)
}
