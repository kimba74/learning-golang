package exercises

import (
	"fmt"
	"learning-golang/utils"
)

func fibonacci() func() int {
	x, y, z := 1, 0, 0
	return func() int {
		z, y, x = y, x, x+y
		return z
	}
}

// Exercise4 runs the corresponding exercise
func Exercise4() {
	utils.ExerciseHeader(4)
	fib := fibonacci()
	fmt.Print("Fibonacci: ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", fib())
	}
	fmt.Println()
}
