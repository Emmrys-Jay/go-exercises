/*

Take the STRUCT “person” in the previous exercise and add a field called “favFood” that stores a slice of string;
for the variable “p1” use a composite literal to add values to the field favFood;
print out the values in favFood;
also print out the values for “favFood” by using a for range loop

*/

package main

import "fmt"

type person struct {
	fname   string
	lname   string
	favFood []string
}

func main() {
	p1 := person{
		"Ade",
		"Tokunbo",
		[]string{"Corn-Flakes", "Rice", "Beans", "Pepperoni", "Sharwama"},
	}

	fmt.Println(p1.favFood)

	for _, food := range p1.favFood {
		fmt.Printf("%s \n", food)
	}
}
