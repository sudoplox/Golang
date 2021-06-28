package main

import (
	"fmt"
	"net/http"

)


func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
	}
	c := make(chan string) //create a channel

	for _,link := range links {
		go checkLink(link, c)
		// concurrency is not parallelism
		// put go keyword infront of function calls

	}
	// fmt.Println( <- c ) // wait for value to be sent into the channel, when we get then log it immediately
	
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string){
	_,err := http.Get(link)
	if err != nil {
		fmt.Println(link,"might be down!")
		c <- "might be down!" // send string to the channel c
		return
	}
	fmt.Println(link,"is up!")
	c <- "Yup it's up!"	// send string to the channel c
}