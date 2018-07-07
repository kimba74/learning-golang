package lessons

import (
	"fmt"
	"learning-golang/utils"
	"math"
)

func function1() {
	fmt.Println("This is function1()")
}

func function2(num int, val string) {
	fmt.Printf("This is function2(%d, \"%s\")\n", num, val)
}

func function3(num int, val string) string {
	return fmt.Sprintf("%d) %s\n", num, val)
}

// powerOf is a closure function
func powerOf(exp int) func(num int) int {
	exponent := exp
	return func(n int) int {
		return int(math.Pow(float64(n), float64(exponent)))
	}
}

// Lesson11 runs the corresponding lesson
func Lesson11() {
	utils.LessonHeader(11)
	function1()
	utils.SectionDivider()
	function2(14, "Hello World")
	utils.SectionDivider()
	fmt.Printf("function3() says: %s", function3(17, "Some String"))
	utils.SectionDivider()
	fmt.Println("Closure \"powerOf\":")
	po := powerOf(2)
	for c := 0; c < 10; c++ {
		fmt.Printf("   %d^2 = %d\n", c, po(c))
	}
}
