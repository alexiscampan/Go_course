package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// for i, arg := range os.Args {
	// 	fmt.Println("argument", i, "is", arg)
	// }
	flag.Usage = func() {
		usageText := `Usage example [OPTION]
		-s, --s example string argument, default: String
		-i, --i example integer argument, default: Int
		-b, --b example of boolean argument, default: Bool`
		fmt.Fprint(os.Stderr, "%s\n", usageText)
	}
	s := flag.String("s", "Hello world", "string help text")
	i := flag.Int("i", 1, "Int help text")
	b := flag.Bool("b", false, "Bool help text")
	flag.Parse()
	fmt.Println("value of s:", *s)
	fmt.Println("value of i:", *i)
	fmt.Println("value of b:", *b)
}
