/*

Add a method to type “person” with the identifier “walk”.
Have the func return this string:
	<person’s first name> is walking.

Remember, you make a func a method by giving the func a receiver.
	func (r receiver) identifier(parameters) (returns) {
	<code>
	}

To return a string, use fmt.Sprintln. Call the method assigning the returned string to a
variable with the identifier “s”. Print out “s”.

*/

package main

import "fmt"

type person struct {
	fname   string
	lname   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fname, "is walking")
}

func main() {
	p1 := person{
		"Ade",
		"Tokunbo",
		[]string{"Corn-Flakes", "Rice", "Beans", "Pepperoni", "Sharwama"},
	}

	str := p1.walk()
	fmt.Println(str)
}
