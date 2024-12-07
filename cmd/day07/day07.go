package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func concat(orig int, toAdd int) int {
	if orig == 0 {
		return 0
	}
	digits := toAdd
	for digits != 0 {
		orig *= 10
		digits /= 10
	}
	return orig + toAdd
}

func solve(target int, current int, values []int) bool {
	if current > target {
		return false
	} else if len(values) == 0 {
		return target == current
	}

	mul := current * values[0]
	add := current + values[0]
	return solve(target, mul, values[1:]) || solve(target, add, values[1:])
}

func solve2(target int, current int, values []int) bool {
	if current > target {
		return false
	} else if len(values) == 0 {
		return target == current
	}

	mul := current * values[0]
	add := current + values[0]
	cat := concat(current, values[0])

	return solve2(target, mul, values[1:]) || solve2(target, add, values[1:]) || solve2(target, cat, values[1:])
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
		if solve(target, 0, values) {
			p1 += target
			p2 += target
		} else if solve2(target, 0, values) {
			p2 += target
		}
	}

	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", p2)
}
