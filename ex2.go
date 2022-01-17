/*

Initialize a MAP using a composite literal where the key is a string and the value is an int;
Print out the map;
range over the map printing out just the key;
range over the map, printing out both the key and the value;

*/
package main

import "fmt"

func main() {
	myMap := map[string]int{"One": 1, "Two": 2, "Three": 3, "Four": 4, "Five": 5}

	fmt.Println(myMap)

	for str, _ := range myMap {
		fmt.Printf("%s ", str)
	}
	fmt.Printf("\n")

	for str, val := range myMap {
		fmt.Printf("%s: %d\n", str, val)
	}
}
