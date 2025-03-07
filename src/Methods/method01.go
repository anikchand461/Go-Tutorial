package main

import "fmt"

/*
The methods for a type are defined at the PACKAGE LEVEL BLOCK

Method declarations look just like function declarations, with one addition: the RECEIVER SPECIFICATION.

The RECEIVER appears between the keyword 'func' and the NAME of the method. Just like
all other variable declarations, the receiver name appears before the type

Just like functions, METHOD NAMES cannot be overloaded. You can use the same method names for
different types, but you can’t use the same name for two different methods on the same type .
*/

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string { // Tells that this method is a property of an object of this struct
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func main() {

	p := Person{
		FirstName: "Fred",
		LastName:  "Fredson",
		Age:       52,
	}
	output := p.String()
	fmt.Println(output)

}
