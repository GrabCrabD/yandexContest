package main

import (
	"fmt"
	"sort"
)

var N, K int

func main() {

	var s int
	resultMap := make(map[int]int)

	arr, sortedArr := myScanner()
	sort.Ints(sortedArr)
	s = sum(sortedArr[1:K+1]) - K*sortedArr[0]
	resultMap[sortedArr[0]] = s

	left, right := 0, K

	for i := 1; i < N; i++ {
		left++
		right--
		diff := sortedArr[i] - sortedArr[i-1]
		s = s - diff*right + diff*(left-1)
		for i+right+1 < N {
			l := sortedArr[i] - sortedArr[i-left]
			r := sortedArr[i+right+1] - sortedArr[i]
			if l > r {
				right++
				left--
				s -= (l - r)
			} else {
				break
			}
		}
		resultMap[sortedArr[i]] = s
	}

	for _, numbers := range arr {
		fmt.Printf("%d ", resultMap[numbers])
	}
}

func sum(arr []int) int {
	var resultSum int
	for _, value := range arr {
		resultSum += value
	}
	return resultSum
}

func myScanner() ([]int, []int) {
	fmt.Scan(&N)
	fmt.Scan(&K)
	arr := make([]int, N)
	copyArr := make([]int, N)
	for i := range arr {
		fmt.Scan(&arr[i])
		copyArr[i] = arr[i]
	}
	return arr, copyArr
}
