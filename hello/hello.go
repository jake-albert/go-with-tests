package main

import "fmt"

const (
	koreanLanguage   = "Korean"
	georgianLanguage = "Georgian"

	englishHello  = "Hello"
	koreanHello   = "안녕"
	georgianHello = "გამარჯობა"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s, %s", getLocalizedHello(language), name)
}

func getLocalizedHello(language string) (hello string) {
	switch language {
	case koreanLanguage:
		hello = koreanHello
	case georgianLanguage:
		hello = georgianHello
	default:
		hello = englishHello
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
