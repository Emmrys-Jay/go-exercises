/*

Using the vehicle, truck, and sedan structs: using a composite literal, create a value of
type truck and assign values to the fields;
Using a composite literal, create a value of type sedan and assign values to the fields.
Print out each of these values. Print out a single field from each of these values.

*/

package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	v         vehicle
	fourWheel bool
}

type sedan struct {
	v      vehicle
	luxury bool
}

func main() {

	truckVal := truck{
		v: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: false,
	}

	sedanVal := sedan{
		v: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(truckVal)
	fmt.Println(sedanVal)
	fmt.Println(truckVal.v)
	fmt.Println(truckVal.fourWheel)
	fmt.Println(sedanVal.v)
	fmt.Println(sedanVal.luxury)

}
