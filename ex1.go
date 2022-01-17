/*

Initialize a SLICE of int using a composite literal;
print out the slice;
range over the slice, printing out just the index;
range over the slice printing out both the index and the value

*/

package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println(mySlice)

	for i, _ := range mySlice {
		fmt.Printf("%d ", i)
	}

	fmt.Printf("\n")

	for i, val := range mySlice {
		fmt.Printf("%d: %d\n", i, val)
	}
}
