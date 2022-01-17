/*

Create a new type called “transportation”. The underlying type of this new type is
interface. An interface defines functionality. Said another way, an interface defines
behavior. For this interface, any other type that has a method with this signature …
transportationDevice() string
… will automatically, implicitly implement this interface. Create a func called “report” that
takes a value of type “transportation” as an argument. The func should print the string
returned by “transportationDevice()” being called for any type that implements the
“transportation” interface. Call “report” passing in a value of type truck. Call “report”
passing in a value of type sedan.

*/

package main

import "fmt"

type transportation interface {
	transportationDevice() string
}

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

func report(t transportation) {
	fmt.Println(t.transportationDevice())
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

	report(truckVal)
	report(sedanVal)
}
