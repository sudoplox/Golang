package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)


func main() {
	resp, err := http.Get("http://google.com") //http part is important
	if err != nil{
		fmt.Println("Error:",err)
		os.Exit(1)
	}

	io.Copy(os.Stdout,resp.Body)	
	// os.Stdout implements Writer interface
	// resp.Body implements Reader interface

	
}