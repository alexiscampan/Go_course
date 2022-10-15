package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "Oh sweet ignition"
	s += " be my fuse"
	fmt.Println(s)

	var buffer bytes.Buffer

	for i := 0; i < 500; i++ {
		buffer.WriteString("z")
	}

	fmt.Println(buffer.String())
	test := "hello"
	fmt.Printf("%b", test[0])

	fmt.Println(strings.ToLower("VERY IMPORTANT MESSAGE"))
	fmt.Println(strings.Index("moon", "oo"))
	fmt.Println(strings.TrimSpace(" I don't need all this space "))

	fmt.Println(strings.ToUpper("test"))
	fmt.Println(strings.ReplaceAll("seaside", "seaside", "bar"))

}
