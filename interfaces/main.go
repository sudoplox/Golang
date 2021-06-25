package main

import (	"fmt")

// bot is interface type, we cant create variable of it
type bot interface{
	//fun_name(arg_type) return_type
	getGreeting() string 
	

	// if you are a type and you have a function called
	// getGreeting and u return a string, 
	// then you can be called a bot
	// thus you can use the function printGreeting too
}

// _bots are concrete type, we can create variable of it
type englishBot struct{}
type spanishBot struct{}


func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// we can omit the variable name since we arent using
func (englishBot) getGreeting() string{
	// very custom logic
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}