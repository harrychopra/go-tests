package hello

import "fmt"

const (
	greetingFmt   = ", %s!"
	enHelloPrefix = "Hello"
	frHelloPrefix = "Bonjour"
	esHelloPrefix = "Hola"
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	greeting := fmt.Sprintf(greetingFmt, name)
	switch lang {
	case "Spanish":
		greeting = esHelloPrefix + greeting
	case "French":
		greeting = frHelloPrefix + greeting
	default:
		greeting = enHelloPrefix + greeting
	}
	return greeting
}
