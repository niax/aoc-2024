package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func isRowSafe(row []int) bool {
	increasing := row[0] < row[1]
	for i := range row {
		if i == 0 {
			continue
		}

		if increasing && !(row[i-1] < row[i]) {
			return false
		}

		if !increasing && !(row[i-1] > row[i]) {
			return false
		}

		delta := abs(row[i] - row[i-1])
		if delta > 3 {
			return false
		}
	}

	return true

}

func main() {
	inputFd, err := os.Open("inputs/02")
	if err != nil {
		panic(err)
	}
	defer inputFd.Close()

	input := bufio.NewScanner(inputFd)
	grid := make([][]int, 0, 1024)
	for input.Scan() {
		line := input.Text()
		nums := strings.Split(line, " ")
		row := make([]int, 0, len(nums))
		for i := range nums {
			v, _ := strconv.Atoi(nums[i])
			row = append(row, v)
		}
		grid = append(grid, row)
	}


	safe := 0
	p2safe := 0
	for _, row := range grid {
		if isRowSafe(row) {
			safe += 1
			p2safe += 1
		} else {
			for i := range row {
				newRow := make([]int, 0, len(row) - 1)
				newRow = append(newRow, row[:i]...)
				newRow = append(newRow, row[i+1:]...)
				if isRowSafe(newRow) {
					p2safe += 1
					break
				}
			}
		}

	}

	fmt.Printf("%d\n%d\n", safe, p2safe)
}
