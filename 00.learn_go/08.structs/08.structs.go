package main

import "fmt"

// data structure,
// collection of properties that are related together
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	john := person{"John", "Smith", contactInfo{"john@aol.com", 1234}}
	jane := person{
		firstName: "Jane",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "jane.doe@gmail.com",
			zipCode: 12345,
		},
	}
	var jenny person

	fmt.Println(john)
	fmt.Println(jane)
	fmt.Println(jenny)
	fmt.Printf("%+v\n", jenny)

	jenny.firstName = "Jenny"
	jenny.lastName = "Bean"
	jenny.contactInfo = contactInfo{
		email:   "jb@gmail.com",
		zipCode: 12345,
	}
	jenny.print()
	jenny.updateName("Julia")
	jenny.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// pass by value
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
