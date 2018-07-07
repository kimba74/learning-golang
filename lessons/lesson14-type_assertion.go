package lessons

import (
	"fmt"
	"learning-golang/utils"
)

// Lesson14 runs the corresponding lesson
func Lesson14() {
	utils.LessonHeader(14)

	typeAssertion()

	utils.SectionDivider()

	typeSwitch("Hello World")
	typeSwitch(256)
	typeSwitch(123.675)
	typeSwitch(true)
}

func typeAssertion() {
	var i interface{}
	var ok bool
	var s string
	var n int

	i = "Hello World"
	fmt.Printf("(%v, %T)\n", i, i)

	s, ok = i.(string)
	fmt.Printf("(%v, %T) => %t\n", s, s, ok)

	n, ok = i.(int)
	fmt.Printf("(%v, %T) => %t\n", n, n, ok)
}

func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("%s is a string\n", v)
	case int:
		fmt.Printf("%d is an int\n", v)
	case float64:
		fmt.Printf("%f is an float64\n", v)
	default:
		fmt.Printf("%v is of unhandled type %T\n", v, v)
	}
}
