/*
Q) How do we run?

$ go run main.go

go build 	-> compiles the files to go source code
go run 		-> compi;es and executes
go fmt 		-> formats all the code in pwd
go install 	-> compiles and installs all pkgs
go get 		-> downloads the raw souce code of external pkg
go test 	-> run any tests in the current project
*/

package main  	//
/*
Q) what is package?

package == project == workspace
grouping of files 

types of packages
	executable
		eg -> main
		file that we can run
	reusable
		codependencies
		reusable logics
		helper functions

main package is 
	for creating executable
	sacred name
	must have main func too
		which will run automatically when program is run
	if we run go build, nothing will happen

*/

import "fmt"  	
/*
Q) What is import and fmt

give package main the functionality of fmt (format) library
eg math, debug, crypto, io, encoding, calculator, uploader
golang.org/pkg
*/

func main(){ 	//func?
	fmt.Println("Hi there")
}
/*
Q) What is func?

function, like def in python
takes arguments in parenthesis
inside function body is the code
*/

/*
Q) Some organisation of code?

package declaration
--then-->
import other packages
--then-->
body - declare functions

*/