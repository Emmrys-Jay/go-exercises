/*

Adding onto this code: Can you assign “g1” to “x”? Why or why not?

*/

/*

g1 cannot be assigned to x because they are not of the same type:

g1 is of type main.gator
x is of type int

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

	x = g1 //ERROR
}
