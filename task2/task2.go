package main

import (
	"fmt"
	"strconv"
)

var N, M int

func main() {
	seatMap, groupWishes := myScanner()
	seatMapMatrix := parseMapToMatrix(seatMap)
	parseGroup(seatMapMatrix, groupWishes)
}

func myScanner() ([]string, []string) {

	fmt.Scan(&N)

	var tmpStr string

	seatMap := []string{}
	for i := 0; i < N; i++ {
		fmt.Scan(&tmpStr)
		seatMap = append(seatMap, tmpStr)
	}

	fmt.Scan(&M)

	groupWishes := []string{}
	for i := 0; i < M*3; i++ {
		fmt.Scan(&tmpStr)
		groupWishes = append(groupWishes, tmpStr)
	}
	return seatMap, groupWishes
}

func getBordersOfSeats(howManyPersons int, leftOrRight string, position string) (int, int) { // определение границ посадки
	var start, end int
	if leftOrRight == "left" && position == "aisle" {
		end = 2
		start = end - howManyPersons + 1
	} else if leftOrRight == "left" {
		start = 0
		end = start + howManyPersons - 1
	} else if position == "aisle" {
		start = 4
		end = start + howManyPersons - 1
	} else {
		end = 6
		start = end - howManyPersons + 1
	}
	return start, end
}

func parseGroup(seatMapMatrix [][]int, groupWishes []string) {

	var status bool
	var resultString string

	for i := 0; i < len(groupWishes); i++ {
		howManyPersons, _ := strconv.Atoi(string(groupWishes[i]))
		leftOrRight := groupWishes[i+1]
		i++
		position := groupWishes[i+1]
		i++

		start, end := getBordersOfSeats(howManyPersons, leftOrRight, position)
		for l := 0; l < N; l++ {
			status = false

			for k := start; k <= end; k++ {
				if seatMapMatrix[l][k] == 1 {
					break
				} else if k == end {
					status = true
				}
			}

			if status == true {
				for marker := start; marker <= end; marker++ {
					seatMapMatrix[l][marker] = 2
				}
				seatMapMatrix = parseToString(seatMapMatrix, l, start, howManyPersons)
				break
			} else if l == N-1 {
				if i != len(groupWishes)-1 {
					fmt.Println("Cannot fulfill passengers requirements")
				} else {
					fmt.Printf("Cannot fulfill passengers requirements")
				}
				break
			}
		}
	}
	fmt.Println(resultString)
}

func parseMapToMatrix(seatMap []string) [][]int {
	seatMapMatrix := make([][]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < 7; j++ {
			if seatMap[i][j] == '.' {
				seatMapMatrix[i] = append(seatMapMatrix[i], 0)
			} else if seatMap[i][j] == '#' {
				seatMapMatrix[i] = append(seatMapMatrix[i], 1)
			} else if seatMap[i][j] == 'X' {
				seatMapMatrix[i] = append(seatMapMatrix[i], 2)
			} else {
				seatMapMatrix[i] = append(seatMapMatrix[i], -1)
			}
		}
	}
	return seatMapMatrix
}

func parseToString(seatMapMatrix [][]int, l int, start int, howManyPersons int) [][]int {

	seatSymbols := [7]string{"A", "B", "C", "_", "D", "E", "F"}
	var resStr string
	rowNumber := strconv.Itoa(l + 1)
	for q := start; q < start+howManyPersons; q++ {
		resStr += rowNumber
		resStr += seatSymbols[q]
		if q != start+howManyPersons-1 {
			resStr += " "
		}
	}
	fmt.Println("Passengers can take seats:", resStr)
	for i := 0; i < N; i++ {
		for j := 0; j < 7; j++ {
			if seatMapMatrix[i][j] == 0 {
				fmt.Print(".")
			} else if seatMapMatrix[i][j] == 1 {
				fmt.Print("#")
			} else if seatMapMatrix[i][j] == 2 {
				fmt.Print("X")
				seatMapMatrix[i][j] = 1
			} else {
				fmt.Print("_")
			}
		}
		fmt.Print("\n")
	}
	return seatMapMatrix
}
