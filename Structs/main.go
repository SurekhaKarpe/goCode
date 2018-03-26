package main

import (
	"fmt"
)

type student struct {
	firstname string
	lastname  string
	age       int
	contact
}

type contact struct {
	city    string
	pincode string
}

func main() {

	Jone := student{
		firstname: "Jone",
		lastname:  "Deo",
		age:       25,
		contact: contact{
			city:    "Pune",
			pincode: "125462",
		},
	}

	Jone.print()

	Jone.updateName("Saurabh")
	Jone.print()

}

func (s student) print() {
	fmt.Printf("%+v", s)
}

func (s *student) updateName(fname string) {
	(*s).firstname = fname
}
