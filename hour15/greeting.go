package testingExample

import (
	"bytes"
	"strings"
)

func translate(s string) string {
	switch s {
	case "en-US":
		return "Hello"
	case "fr-FR":
		return "Bonjour"
	case "it-IT":
		return "Ciao"
	default:
		return "Hello"
	}
}

func Greeting(name, locale string) string {
	salutation := translate(locale)
	return (salutation + " " + name)
}

func StringFromAssignement(j int) string {
	var s string
	for i := 0; i < j; i++ {
		s += "a"
	}
	return s
}

func StringFromAppendJoin(j int) string {
	s := []string{}
	for i := 0; i < j; i++ {
		s = append(s, "a")
	}
	return strings.Join(s, "")
}

func StringFromBuffer(j int) string {
	var buffer bytes.Buffer
	for i := 0; i < j; i++ {
		buffer.WriteString("a")
	}
	return buffer.String()
}
