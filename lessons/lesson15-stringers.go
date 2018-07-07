package lessons

import (
	"fmt"
	"learning-golang/utils"
)

type person15 struct {
	name string
	age  int
}

func (p person15) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

// Lesson15 runs the corresponding lesson
func Lesson15() {
	utils.LessonHeader(15)
	p1 := person15{"Malcom Taylor", 38}
	p2 := person15{"Valerie Marquat", 43}

	fmt.Printf("This is: %s\n", p1)
	fmt.Printf("This is: %s\n", p2)
}
