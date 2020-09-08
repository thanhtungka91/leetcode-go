package main

import "fmt"

type Person struct {
	// Address
	StreetAddress, PostCode, City string

	// Job

	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func (it *PersonBuilder)Build()  *Person{
	return it.person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		person: &Person{},
	}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (ps *PersonBuilder) Lives() *PersonAddressBuilder  {
	return &PersonAddressBuilder{*ps}
}

func (ps *PersonBuilder) Works() *PersonJobBuilder  {
	return &PersonJobBuilder{*ps}
}

func (ps *PersonAddressBuilder)At(streetAdress string)   *PersonAddressBuilder{
	ps.person.StreetAddress = streetAdress
	return ps
}


func (ps *PersonJobBuilder)At(companyAddress string)   *PersonJobBuilder{
	ps.person.CompanyName = companyAddress
	return ps
}

func main() {

	pb := NewPersonBuilder()
	pb.Lives().At("Tokyo Japan").Works().At("Home")

	person := pb.Build()
	fmt.Println(person)
	fmt.Println(pb.person)
}
