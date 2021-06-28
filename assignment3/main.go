package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)


func main() {
	fmt.Println(os.Args[1])
	f,err := ioutil.ReadFile( string(os.Args[1]) )
	if err != nil {
		log.Fatal(err)
   }

  fmt.Println(string(f))
}
