package main

import "fmt"

func max(a [5]int) (err int, val int) {
	if len(a) == 0 {
		err = -1
	}

	if a[0] < 0 {
		err = -2
	}

	val = a[0]

	i := 1
	for i < len(a) {
		if a[i] < 0 {
			err = -2
		}
		if a[i] > val {
			val = a[i]
		}
	}
	return
}

func test(a [5]int) {
	fail, maxval := max(a)

	fmt.Println("Printing array... ")
	fmt.Println(a)

	if fail < 0 {
		fmt.Println("FAIL => error code: ")
		fmt.Println(fail)
	} else {
		fmt.Println("The maximum value in the array is: ")
		fmt.Println(maxval)
	}
}

func main() {
	var a1 [5]int
	a1[0] = 5
	a1[1] = 9
	a1[2] = 3
	a1[3] = 9
	a1[4] = 7

	var a2 [5]int

	var a3 [5]int
	a3[0] = 5
	a3[1] = -9
	a3[2] = 3
	a3[3] = 2
	a3[4] = 6

	test(a1)
	test(a2)
	test(a3)
}
