package lessons

import (
	"fmt"
	"learning-golang/utils"
)

var gv1 int

var gv2 = "A simple String"

var gv3, gv4 float64 = 1.23, 2.34

var gv5, gv6 = "Another String", false

var (
	gv7 = 6.78
	gv8 = 123
)

// Lesson1 runs the corresponding lesson
func Lesson1() {
	utils.LessonHeader(1)

	var lv1 bool

	var lv2, lv3 float64

	var lv4, lv5 int = 234, 567

	var lv6, lv7 = true, 885

	var lv8 float64
	lv8 = 77.89

	fmt.Println("Global Variables:")
	fmt.Printf("   gv1 = %v (%T)\n", gv1, gv1)
	fmt.Printf("   gv2 = %v (%T)\n", gv2, gv2)
	fmt.Printf("   gv3 = %v (%T)\n", gv3, gv3)
	fmt.Printf("   gv4 = %v (%T)\n", gv4, gv4)
	fmt.Printf("   gv5 = %v (%T)\n", gv5, gv5)
	fmt.Printf("   gv6 = %v (%T)\n", gv6, gv6)
	fmt.Printf("   gv7 = %v (%T)\n", gv7, gv7)
	fmt.Printf("   gv8 = %v (%T)\n", gv8, gv8)

	fmt.Println("Function Variables:")
	fmt.Printf("   lv1 = %v (%T)\n", lv1, lv1)
	fmt.Printf("   lv2 = %v (%T)\n", lv2, lv2)
	fmt.Printf("   lv3 = %v (%T)\n", lv3, lv3)
	fmt.Printf("   lv4 = %v (%T)\n", lv4, lv4)
	fmt.Printf("   lv5 = %v (%T)\n", lv5, lv5)
	fmt.Printf("   lv6 = %v (%T)\n", lv6, lv6)
	fmt.Printf("   lv7 = %v (%T)\n", lv7, lv7)
	fmt.Printf("   lv8 = %v (%T)\n", lv8, lv8)
}
