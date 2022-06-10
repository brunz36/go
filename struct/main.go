package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	info := contactInfo{email: "thais@gmail.com", zipCode: 34677}
	thais := person{"Thais", "Ribeiro", info}
	olivia := person{firstName: "Olivia", lastName: "Ribeiro", contact: contactInfo{"olivia@gmail.com", 34677}}

	thais.print()
	olivia.print()

	var bruno person
	bruno.print()
	bruno.firstName = "Bruno"
	bruno.lastName = "Ribeiro"
	bruno.contact = contactInfo{"bruno@gmail.com", 34677}
	bruno.print()
	bruno.updateFirstName("Bubu")
	bruno.print()
}

func (p *person) updateFirstName(name string) {
	(*p).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
