package main

import "fmt"

func max(a [3]int) (val int, err int) {

	if len(a) == 0 {
		err = -1
	}

	if a[0] < 0 {
		err = -2
	}

	val = a[0]

	i := 1
	j := len(a)
	for i != j {
		if a[i] < 0 {
			err = -2
		}
		if a[i] > val {
			val = a[i]
		}
		i++
	}
	return val, err
}

func test(a [3]int) {
	maxval, fail := max(a)

	fmt.Println("Printing array... ")
	fmt.Println(a)

	if fail < 0 {
		fmt.Printf("FAIL => error code: %d\n", fail)
	} else {
		fmt.Printf("The maximum value in the array is: %d\n", maxval)
	}
}

func main() {
	var a1 [3]int
	a1[0] = 5
	a1[1] = 9
	a1[2] = 3

	var a2 [3]int

	var a3 [3]int
	a3[0] = 5
	a3[1] = -9
	a3[2] = 3

	fmt.Println()
	test(a1)
	fmt.Println()
	test(a2)
	fmt.Println()
	test(a3)
}
