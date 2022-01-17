/*

Adding onto this code: Using var, declare an identifier “x” as type int (var x int). Print out
“x”. Print the type of “x” using fmt.Printf(“%T\n”, x)

*/

package main

import "fmt"

type gator int

var g1 gator
var x int

func main() {
	g1 = 45
	x = 9
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
