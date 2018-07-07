package utils

import "fmt"

// ExerciseHeader creates a standard header for exercises
func ExerciseHeader(num int) {
	fmt.Printf("\n--8<------- EXERCISE %d -----------------8<------------------\n\n", num)
}

// LessonHeader creates a standard header for lessons
func LessonHeader(num int) {
	fmt.Printf("\n--8<------- LESSON %d ------------------8<------------------\n\n", num)
}

// SectionDivider creates a devider within a lesson or exercise
func SectionDivider() {
	fmt.Println("\n  ---------------------------- \n")
}
