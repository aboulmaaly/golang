/************************************************
 * Function Returning Another Function          *
 ************************************************
 *
 * We see functions x and y which return another
 * lambda function : func(b int) int
 *
 * func x() (func(b int) int)
 * func y(a int) (func(b int) int)
 *
 * x() takes no parameters
 * y() takes an int as parameter
 *
 ************************************************
 * Outputs :                                    *
 ************************************************
 *
 * Call x() for 3 gives : 5
 * Call y(2) for 3 gives : 5
 *
 */

package main

import "fmt"

func main() {
	// Make a func x(), give it a name p, and call it.
	p := x()
	fmt.Printf("Call x() for 3 gives : %v\n", p(3))

	// Make a special func y(a int), a gets 3.
	p2 := y(2)
	fmt.Printf("Call y(2) for 3 gives : %v\n", p2(3))
}

func x() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func y(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
