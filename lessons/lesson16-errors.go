package lessons

import (
	"fmt"
	"learning-golang/utils"
	"time"
)

// myError is a struct that describes
// the error thrown by the functions in this lesson
type myError struct {
	when time.Time
	what string
}

// Error is an implementation of an interface provided by the
// fmt package. It does the same for error values as Stringer
// does for structs.
func (e myError) Error() string {
	return fmt.Sprintf("[%v, \"%s\"]", e.when.Format(time.RFC3339), e.what)
}

// doSomething will take the value parameter and does a modulo operation
// on it using the provided module parameter and return the result.
// If the result is 0 the function will also return an error value.
func doSomething(value, modulo int) (int, *myError) {
	result := value % modulo
	if result != 0 {
		return result, nil
	}
	return result, &myError{
		time.Now(),
		fmt.Sprintf("modulo result of %d and %d is 0!", value, modulo),
	}
}

// Lesson16 runs the corresponding lesson
func Lesson16() {
	utils.LessonHeader(16)
	for i := 0; i < 10; i++ {
		val, err := doSomething(i, 3)
		if err != nil {
			fmt.Printf("%d mod 3 = %d %s\n", i, val, err)
		} else {
			fmt.Printf("%d mod 3 = %d\n", i, val)
		}
	}
}
