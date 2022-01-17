/*

Give a method to both the “truck” and “sedan” types with the following signature
transportationDevice() string
Have each func return a string saying what they do. Create a value of type truck and
populate the fields. Create a value of type sedan and populate the fields. Call the
method for each value.

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

func (t truck) transportationDevice() string {
	return fmt.Sprintln("I am a", t.v.color, "truck with", t.v.doors, "doors and many wheels")
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintln("I am a", s.v.color, "sedan with", s.v.doors, "doors and luxury")
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

	str := sedanVal.transportationDevice()
	fmt.Println(str)
	str = truckVal.transportationDevice()
	fmt.Println(str)

}
