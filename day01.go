package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputFd, err := os.Open("day01.txt")
	if err != nil {
		panic(err)
	}
	defer inputFd.Close()

	input := bufio.NewScanner(inputFd)
	left := make([]int, 0, 1024)
	right := make([]int, 0, 1024)
	for input.Scan() {
		line := input.Text()
		parts := strings.Split(line, "   ")
		l, _ := strconv.Atoi(parts[0])
		left = append(left, l)
		r, _ := strconv.Atoi(parts[1])
		right = append(right, r)
	}
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	diffSum := 0 
	for i := range left {
		diff := max(left[i], right[i]) - min(left[i], right[i])
		diffSum += diff
	}
	fmt.Printf("%d\n", diffSum)

	similarityScore := 0
	lIdx := 0
	rIdx := 0
	for lIdx < len(left) {
		lVal := left[lIdx]
		for rIdx < len(right) {
			rVal := right[rIdx]
			if lVal == rVal {
				similarityScore += lVal
			} else if rVal > lVal {
				break
			}
			rIdx++
		}
		lIdx++
	}

	fmt.Printf("%d\n", similarityScore)
}
