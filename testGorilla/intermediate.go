/*

Write a function that accepts two numbers n, and k
and generates a series of n numbers where each number is represented
by x in 1 <= x <= n and are TVstations

For every addition of k, the TV station is deleted, and so on. On getting
to the last TV station, the channels move back to the first channel.

The function will return the only channel to remain after all channels have
been deleted.

******** Example 1

For Example for an input of n=5 and k=2
ch1+2 = ch3 (ch3 is deleted)
channels left are 1, 2, 4, 5

ch4+2 = ch5, then back to ch1 (ch1 is deleted)
channels left are 2, 4, 5

ch2+2 = ch5 (ch5 is deleted)
channels left are 2, 4

ch2+2 = ch2 (ch2 is deleted)
channels left is 4

Therefore, the function returns 4

******** Example 2

For Example for an input of n=5 and k=3
ch1+3 = ch4 (ch4 is deleted)
channels left are 1, 2, 3, 5

ch3+3 = ch5, then back to ch2 (ch2 is deleted)
channels left are 1, 3, 5

ch3+3 = ch3 (ch3 is deleted)
channels left are 1, 5

ch5+3 = ch5 (ch5 is deleted)
channels left is 1

Therefore, the function returns 1
*/

package main

import "fmt"

func RandomFunc2(n, k int) int {
	// Populate array
	TVstations := []int{}
	for i := 0; i < n; i++ {
		TVstations = append(TVstations, i+1)
	}

	i := 0
	for len(TVstations) != 1 {
		if (len(TVstations) - (i + 1)) < k {
			i = (k - ((len(TVstations) - 1) - i)) - 1
		} else {
			i += k
		}

		TVstations = append(TVstations[:i], TVstations[i+1:]...)

		if len(TVstations) == i {
			i = 0
		}
	}

	return TVstations[0]
}

func main() {
	fmt.Println(RandomFunc2(5, 3))
}
