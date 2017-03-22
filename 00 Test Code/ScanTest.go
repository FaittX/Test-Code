package main

import "fmt"

var name string

func main() {
	fmt.Print("Enter your name: ")
	fmt.Scan(&name )
	fmt.Print("Hello," , name)

}
