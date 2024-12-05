package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/niax/aoc-2024/internal/collections"
)

func main() {
	inputFd, err := os.Open("inputs/05")
	if err != nil {
		panic(err)
	}
	defer inputFd.Close()

	input := bufio.NewScanner(inputFd)
	rules := make(map[int]collections.Set[int], 1024)
	for input.Scan() {
		line := input.Text()
		if line == "" {
			break
		}
		s := strings.SplitN(line, "|", 2)
		a, _ := strconv.Atoi(s[0])
		b, _ := strconv.Atoi(s[1])
		_, ok := rules[a]
		if !ok {
			rules[a] = collections.NewSet[int]()
		}
		rules[a].Add(b)
	}

	cmp := func(a, b int) int {
		s, ok := rules[a]
		if ok {
			if s.Contains(b) {
				return -1
			}
		}
		return 0
	}

	p1 := 0
	p2 := 0
	for input.Scan() {
		line := input.Text()
		s := strings.Split(line, ",")

		lineNums := make([]int, len(s))
		for i := range s {
			page, _ := strconv.Atoi(s[i])
			lineNums[i] = page
		}

		toUpdate := &p1
		if slices.IsSortedFunc(lineNums, cmp) {
			toUpdate = &p1
		} else {
			toUpdate = &p2
			slices.SortFunc(lineNums, cmp)
		}
		midpoint := lineNums[len(lineNums)/2]
		*toUpdate += midpoint
	}

	fmt.Printf("%d\n%d\n", p1, p2)
}
