package main

import (	"fmt")

type englishBot struct{}
type spanishBot struct{}


func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	// printGreeting(sb)

}

// we can omit the variable name since we arent using
func (englishBot) getGreeting() string{
	// veru custom logic
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

// no function overloading?
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}
// func printGreeting(sb spanishBot){
// 	fmt.Println(sb.getGreeting())
// }

