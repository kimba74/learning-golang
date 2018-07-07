package lessons

import (
	"fmt"
	"learning-golang/utils"
)

// Name is a struct that captures a default western name
type name struct {
	first, middle, last string
}

// ToString will print out a Name struct
func (n name) toString() string {
	return "Name" +
		"\n  First  : " + n.first +
		"\n  Middle : " + n.middle +
		"\n  Last   : " + n.last
}

// Person is a struct describing a person
type person struct {
	name name
}

// ToString will print out a Person struct
func (p person) toString() string {
	return "Person" +
		"\n  " + p.name.last + ", " + p.name.first + " " + p.name.middle
}

// Lesson12 runs the corresponding lesson
func Lesson12() {
	utils.LessonHeader(12)
	p := person{}
	n := name{first: "Steffen", middle: "Lars", last: "Krause"}
	p.name = n

	fmt.Println(p.toString())
	fmt.Println(n.toString())
}
