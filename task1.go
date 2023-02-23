package main

import "fmt"

func main() {

	var max, result = 0, 0
	var n, a, min = 1, 1, 1
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a)

		if i == 0 {
			max = a
			min = a
		}

		if max > a {
			result = -1
			break
		}

		if min > a {
			min = a
		}

		max = a
	}
	if result == 0 {
		fmt.Println(max - min)
	} else {
		fmt.Println(result)
	}
}
