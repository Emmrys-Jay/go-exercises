/*

Using the short declaration operator, create a variable with the identifier “s” and assign
“i'm sorry dave i can't do that” to “s”.
i.
Print “s”.
ii.
Print “s” converted to a slice of byte.
iii.
Print “s” converted to a slice of byte and then converted back to a string.
iv.
Using slicing, print just “i’m sorry dave”
v.
Using slicing, print just “dave i can’t”
vi.
Using slicing, print just “can’t do that”
vii.
print every letter of “s” with one rune (character) on each line

*/

package main

import (
	"fmt"
)

func main() {
	s := "i'm sorry dave i can't do that"
	fmt.Println(s)

	str := []byte(s)
	fmt.Println(str)

	fmt.Println(string(str))

	fmt.Println(string(str[0:14]))

	fmt.Println(string(str[10:22]))

	fmt.Println(string(str[17:]))

	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}
}
