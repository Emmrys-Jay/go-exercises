/*

Now add a method to type gator with this signature ...
greeting()
… and have it print “Hello, I am a gator”. Create a value of type gator. Call the greeting()
method from that value.

*/

package main

import "fmt"

type gator int

func (g gator) greeting() {
	fmt.Printf("Hello, i am %T\n", g)
}

func main() {

	var g1 gator
	g1.greeting()

}
