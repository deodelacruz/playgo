/* Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.

s =½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement. Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.
fmt.Println(fn(3))
And I can use the following statement to print the displacement after 5 seconds.
fmt.Println(fn(5)) */

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hi. Please type in a value for acceleration.")
	input, _ := reader.ReadString('\n')
	a, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	fmt.Println("Initial velocity:")
	input, _ = reader.ReadString('\n')
	vo, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	fmt.Println("Initial displacement:")
	input, _ = reader.ReadString('\n')
	so, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	fmt.Println("Elapsed time:")
	input, _ = reader.ReadString('\n')
	t, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	func1 := GenDisplaceFn(a, vo, so)
	fmt.Printf("Ok. The computed displacement is: %v \n", func1(t))
}

func GenDisplaceFn(a float64, vo float64, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (0.5 * a * math.Pow(t, 2)) + (vo * t) + so
	}
	return fn
}
