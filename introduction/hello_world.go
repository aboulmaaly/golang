/********************************************************************
 * package main :
 *		packages are used to provide code
 *		compartmentalisation and reusability.
 *
 * import "fmt" :
 *		The fmt package is imported and it will be used inside
 *		the main function to print text to the stadard output.
 *		A package can also be given another name (alias)
 *		like : import fm "fmt".
 *		This alias is then will be used in the code above
 *		like : fm.Println("Hello world")
 *
 * func main() :
 *		The main is a special function.
 *		The program execution starts from the main function.
 *		The main function should always reside in the main package.
 *
 * fmt.Println("Hello world") :
 *		The Println function of the fmt package is used to write
 *		text to the standard output.
 *
 */
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
