package main

import "fmt"

/*
*
the interface is a custom type that is used to specify a set of one or more method signatures and the interface is abstract,
so you are not allowed to create an instance of the interface. But you are allowed to create a variable of an interface type
and this variable can be assigned with a concrete type value that has the methods the interface requires.
Or in other words, the interface is a collection of methods as well as it is a custom type.
*
*/

// this is example

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MethodPerson interface {
	SetName(name string) (result string)
}

func (m Person) SetName(name string) string {
	m.Name = name
	return m.Name
}

func main() {
	var person Person

	result := person.SetName("Farhani")

	fmt.Println(result)
}
