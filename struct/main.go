package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Printf("%+v", alex)

	var anon person
	anon.lastName = " 4n0n"
	fmt.Printf("%+v", anon)
}
