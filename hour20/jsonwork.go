package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name    string   `json:"name,omitempty"`
	Age     int      `json:"age,omitempty"`
	Hobbies []string `json:"hobbies,omitempty"`
}

type Switch struct {
	On bool `json:"on"`
}

func main() {
	// hobbies := []string{"Cycling ", "Cheese", "Techno"}
	// p := Person{
	// 	Name:    "George",
	// 	Age:     40,
	// 	Hobbies: hobbies,
	// }
	// fmt.Printf("%+v\n", p)
	// jsonByteData, err := json.Marshal(p)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// jsonStringData := string(jsonByteData)
	// fmt.Println(jsonStringData)

	// jsonStringData := `{"name":"George","age":40,"hobbies":["Cycling", "Cheese","Techno"]}`
	// jsonByteData := []byte(jsonStringData)
	// p := Person{}
	// err := json.Unmarshal(jsonByteData, &p)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", p)

	// jsonStringData := `{"on":"true"}`
	// jsonByteData := []byte(jsonStringData)
	// s := Switch{}
	// err := json.Unmarshal(jsonByteData, &s)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", s)

	var u User
	res, err := http.Get("https://api.github.com/users/shapeshed")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", u)

}
