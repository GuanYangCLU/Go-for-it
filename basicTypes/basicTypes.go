package main

import "fmt"

//package scope
type myType int

type person struct {
	fname string // if want use outside capitalize first?
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Hello"`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Hello"`)
}

// polymorphism
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// block scope
	var mt myType
	mt = 3
	fmt.Println(mt)

	xi := []int{2, 3, 4, 5} // Slice, in same line can without last comma
	fmt.Println(xi)

	m := map[string]int{
		"apple": 1,
		"bible": 2, // must with comma since in different line
	}
	fmt.Println(m)

	p1 := person{
		"John",
		"Doe",
	}
	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person{
			"Victor",
			"Lee",
		},
		true,
	}
	sa1.speak()
	sa1.person.speak()

	saySomething(p1)
	saySomething(sa1)
}
