package main

import (
	"fmt"
)

func main() {
	arr := []int{0,1,2,3,4,5,6,7,8,9,10}
	for _,a := range arr{
		if a%2==0{
			fmt.Println(a, "is even")
		} else{
			fmt.Println(a, "is odd")}
	}
}
