package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	// needle := "(?i)chocolate"
	// haystack := "Chocolate is my favorite!"
	// match, err := regexp.MatchString(needle, haystack)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(match)

	// re := regexp.MustCompile("^[a-zA-Z0-9]{5,12}")
	// fmt.Println(re.MatchString("slimshady99"))
	// fmt.Println(re.MatchString("!asdf£33£3"))
	// fmt.Println(re.MatchString("roger"))
	// fmt.Println(re.MatchString("iamthebestofusethisappevaaa"))

	// usernames := [4]string{
	// 	"slimshady99",
	// 	"!asdf£33£3",
	// 	"roger",
	// 	"Iamthebestuserofthisappevaaaar",
	// }
	// re := regexp.MustCompile("^[a-zA-Z0-9]{5,12}")
	// an := regexp.MustCompile(":^alnum:")

	// for _, username := range usernames {
	// 	if len(username) > 12 {
	// 		username = username[:12]
	// 		fmt.Printf("trimmed username to %v\n", username)
	// 	}
	// 	if !re.MatchString(username) {
	// 		username = an.ReplaceAllString(username, "x")
	// 		fmt.Printf("rewrote username to %v\n", username)
	// 	}
	// }
	resp, err := http.Get("https://petition.parliament.uk/petitions")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	src := string(body)
	re := regexp.MustCompile(`\\<h3\\>.*\\</h3\\>`)
	titles := re.FindAllString(src, -1)
	for _, title := range titles {
		fmt.Println(title)
	}
}
