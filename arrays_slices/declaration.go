/************************************************
 * Declaration and Initialization               *
 ************************************************
 *
 * An arrar is a numbered and fixed length sequence
 * of data items (elements) of the same single type.
 *
 * This type can be anything from primitve types
 * like itegers, strings ...
 *
 * The length must be a constant expression, that must
 * evaluated to a non-negative integer value.
 *
 * It's part of the type of the array, so arrays
 * declared as [5]int and [10]int differ in type.
 *
 * The items can be accessed and changed through their
 * index (the position), the index starts from 0,
 * the 2 nd index 1, etc.
 *
 * The number of items, also called length (len) or size
 * of the array, is fixed and must be given when declaring
 * the array; the maximum array length is 2 Gb.
 *
 * the format of the declaration is :
 *		var identifier [len]type
 * ex :
 *		var arr1 [5]int
 *
 * When declaring an array, each item is automatically
 * initialized with the default zero value of the type,
 * in arr1 all items default to 0.
 *
 * The length of arr1 len(arr1) is 5
 * The index ranges from 0 to len(arr1)-1
 * The first element is given by arr1[0]
 * The last element is given by arr1[len(arr1)-1]
 *
 * Assigning a value to an array item at the index i
 * is done by arr[i] = value
 *
 */
// A basic example :

package main

import "fmt"

func main() {
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		// double each index of the arr1
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		// print each index of the arr1
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}
}
