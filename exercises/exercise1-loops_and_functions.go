package exercises

import (
	"fmt"
	"learning-golang/utils"
	"math"
)

func sqrt(x float64) float64 {
	z, o := 1.0, 0.0
	for c := 15; z-o != 0 && c > 0; c-- {
		o = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(o, z, z-o)
	}
	return z
}

// Exercise1 runs the corresponding exercise
func Exercise1() {
	utils.ExerciseHeader(1)
	fmt.Println(sqrt(12))
	fmt.Println(math.Sqrt(12))
}
