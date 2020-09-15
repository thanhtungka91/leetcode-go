package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
	City string
}

type By func(p1, p2 *Person) bool

// main sort after builder type By

func (by By) Sort(people []Person) {
	// using peple sorter to sort this
	pp := &peopleSorter{
		people: people,
		by:     by,
	}

	sort.Sort(pp) // sort(*pointer)
}

type peopleSorter struct {
	people []Person
	by     By
}

// 3 functions for sort interface {
// Len return int
// Swap no return
// Less  return bool() is less or not
// } /

func (p *peopleSorter) Len() int {
	return len(p.people)
}

func (p *peopleSorter) Swap(i, j int) {
	p.people[i], p.people[j] = p.people[j], p.people[i]
}

func (p *peopleSorter) Less(i, j int) bool {
	return p.by(&p.people[i], &p.people[j]) // day gia tri vao
}

func main() {

	people := []Person{
		{"Tung", 20, "Hnoi"},
		{"Yamada", 50, "Tokyo"},
		{"Alex", 40, "Sanfracisco"},
	}

	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}

	age := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}

	By(name).Sort(people)

	fmt.Println("sort by Name", people)

	By(age).Sort(people)

	fmt.Println("sort by Age", people)

}
