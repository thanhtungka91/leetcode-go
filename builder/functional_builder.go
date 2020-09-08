package main

import "fmt"

type Person2 struct {
	name, position string
}

type personMod func(*Person2)

type PersonBuilder2 struct {
	actions []personMod
}


func (b *PersonBuilder2) Called(name string) *PersonBuilder2 {
	b.actions = append(b.actions, func(person2 *Person2) {
		person2.name = name
	})
	return b
}

func (b *PersonBuilder2) Build() *Person2  {
	p:= Person2{}

	for _, a := range b.actions{
		a(&p)
	}

	return &p
}

func (b *PersonBuilder2)WroksAsA(position string) *PersonBuilder2  {
	b.actions = append(b.actions, func(person2 *Person2) {
		person2.position = position
	})
	return b
}

func main() {
	b:= PersonBuilder2{}
	p:= b.Called("Tesstt").WroksAsA("engineer").Build()

	fmt.Println(*p)
}


