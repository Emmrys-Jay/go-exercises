/*

Write a program that accepst an input number
then returns the index of that number in a series n1, n2, n3,...
where the series begins with n1=1, n2=2, n3=3, and n3=4 and
subsquent numbers are a product of the previous four numbers

*/

package main

import "fmt"

func RandomFunc(num int) int {
	var numbers = []int{1, 2, 3, 4}

	if num < 1 {
		return 0
	}

	if num <= 4 {
		return numbers[num-1]
	}

	for i := 4; i < num; i++ {
		tnum := numbers[i-1] + numbers[i-2] + numbers[i-3] + numbers[i-4]
		numbers = append(numbers, tnum)
	}

	return numbers[num-1]

}

func main() {
	fmt.Println(RandomFunc(7))
}
