package main

import (
	"fmt"
	"net/http"
	"os"
)


func main() {
	resp, err := http.Get("http://google.com") //http part is important
	if err != nil{
		fmt.Println("Error:",err)
		os.Exit(1)
	}
	bs := make([]byte,99999) //number of initialised values
	resp.Body.Read(bs)
	fmt.Println(string(bs))
	
}