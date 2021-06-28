package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com") //http part is important
	if err != nil{
		fmt.Println("Error:",err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw,resp.Body)	

	// os.Stdout implements Writer interface
	// resp.Body implements Reader interface
}

func (logWriter) Write(bs []byte) (int, error){
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:",len(bs))
	return len(bs),nil

}