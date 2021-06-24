package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email : "jim@gmail.com",
			zipCode : 110021,
		},
	}
	jimPointer := &jim 	//give the memory address of the variabel
	jimPointer.updateName("jimmy")
	jim.print()
}
// turn address into value 		with 	*address
// turn value 	into address 	with 	&value

/*
*person -> this is a type description, it means we are working with a pointer
*pointerToPerson -> this is an operator, it means get the value at that address
*/
func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName // give the value at that memory address
}


func (p person) print(){
	fmt.Printf("%+v",p)
}