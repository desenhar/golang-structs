package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
