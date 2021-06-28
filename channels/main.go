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
		go func (link string)  { // to fix the prob
			time.Sleep(time.Second)
			checkLink(link,c)  // prob: l variable from main routine but used in child routine. pass l value instead
		}(l) // to fix the prob
	}
}

func checkLink(link string, c chan string){
	_,err := http.Get(link)

	if err != nil {
		fmt.Println(link,"might be down!")
		c <- link // send string to the channel c
		return
	}

	fmt.Println(link,"is up!")
	c <- link	// send string to the channel c
}