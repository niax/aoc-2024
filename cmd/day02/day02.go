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
	if len(row) < 2 {
		return true
	}
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

func part2(current, remaining []int, removed bool) bool {
	if len(remaining) == 0 {
		return true
	}

	next := make([]int, 0, len(current)+1)
	next = append(next, current...)
	next = append(next, remaining[0])
	remaining = remaining[1:]

	if isRowSafe(next) {
		if removed {
			return part2(next, remaining, true)
		} else {
			return part2(next, remaining, false) || part2(current, remaining, true)
		}
	} else {
		if removed {
			return false
		} else {
			return part2(current, remaining, true)
		}
	}
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
			if part2([]int{}, row, false) {
				p2safe += 1
			}
		}

	}

	fmt.Printf("%d\n%d\n", safe, p2safe)
}
