package main

import (
	"fmt"
	"net/http"
	"time"
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
	
	for l := range c {
		go checkLink( l,c)
		
	}
}

func checkLink(link string, c chan string){
	_,err := http.Get(link)
	
	time.Sleep(time.Second)

	if err != nil {
		fmt.Println(link,"might be down!")
		c <- link // send string to the channel c
		return
	}
	
	fmt.Println(link,"is up!")
	c <- link	// send string to the channel c
}