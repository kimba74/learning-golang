package lessons

import (
	"fmt"
	"learning-golang/utils"
)

type empty interface{}

type printer interface {
	print()
	decorate(decorator string)
	decoratePlus(prefix, suffix string)
}

type banner struct {
	title    string
	subtitle string
}

func (b banner) print() {
	fmt.Printf("%s (%s)\n", b.title, b.subtitle)
}

func (b banner) decorate(decorator string) {
	fmt.Printf("%s %s (%s) %s\n", decorator, b.title, b.subtitle, decorator)
}

func (b banner) decoratePlus(prefix, suffix string) {
	fmt.Printf("%s %s (%s) %s\n", prefix, b.title, b.subtitle, suffix)
}

// Lesson13 runs the corresponding lesson
func Lesson13() {
	utils.LessonHeader(13)
	var i printer
	b := banner{title: "GOLANG", subtitle: "how the cloud is done"}

	i = b
	i.print()
	i.decorate("*")
	i.decoratePlus("+", "%")

	utils.SectionDivider()

	var e empty
	e = "Hello World"
	fmt.Printf("(%v, %T)\n", e, e)
	fmt.Printf("(%v, %T)\n", i, i)
}
