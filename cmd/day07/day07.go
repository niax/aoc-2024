package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func digitMask(i int) int {
	if i >= 1000000 {
		panic(i)
	} else if i >= 100000 {
		return 1000000
	} else if i >= 10000 {
		return 100000
	} else if i >= 1000 {
		return 10000
	} else if i >= 100 {
		return 1000
	} else if i >= 10 {
		return 100
	} else {
		return 10
	}
}

func backSolve(target int, values []int) bool {
	if target == 0 && len(values) == 0 {
		return true
	} else if len(values) == 0 {
		return false
	}

	lastVal := values[len(values)-1]
	remaining := values[:len(values)-1]

	if target%lastVal == 0 {
		if backSolve(target/lastVal, remaining) {
			return true
		}
	}

	return backSolve(target-lastVal, remaining)
}

func backSolve2(target int, values []int) bool {
	if target == 0 && len(values) == 0 {
		return true
	} else if len(values) == 0 {
		return false
	}

	lastVal := values[len(values)-1]
	remaining := values[:len(values)-1]
	if target%lastVal == 0 {
		if backSolve2(target/lastVal, remaining) {
			return true
		}
	}
	mask := digitMask(lastVal)
	if target%mask == lastVal {
		if backSolve2(target/mask, remaining) {
			return true
		}
	}

	return backSolve2(target-lastVal, remaining)
}

func main() {
	inputFd, err := os.Open("inputs/07")
	if err != nil {
		panic(err)
	}
	defer inputFd.Close()

	input := bufio.NewScanner(inputFd)
	p1 := 0
	p2 := 0
	for input.Scan() {
		line := input.Text()
		parts := strings.SplitN(line, ": ", 2)
		target, _ := strconv.Atoi(parts[0])
		valueStrs := strings.Split(parts[1], " ")
		values := make([]int, len(valueStrs))
		for i := range valueStrs {
			values[i], _ = strconv.Atoi(valueStrs[i])
		}
		if backSolve(target, values) {
			p1 += target
			p2 += target
		} else if backSolve2(target, values) {
			p2 += target
		}
	}

	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", p2)
}
