package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func shittyLog10(i int) int {
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func main() {
	inputFd, err := os.Open("inputs/11")
	if err != nil {
		panic(err)
	}
	defer inputFd.Close()

	counter := make(map[int]int, 1024)
	input := bufio.NewScanner(inputFd)
	for input.Scan() {
		line := input.Text()
		parts := strings.Split(line, " ")
		for i := range parts {
			v, _ := strconv.Atoi(parts[i])
			counter[v]++
		}
	}

	p1 := 0
	p2 := 0

	nextCounter := make(map[int]int, 1024)
	for step := 0; step < 75; step++ {
		if step == 25 {
			for _, count := range counter {
				p1 += count
			}
		}

		for v, count := range counter {
			if v == 0 {
				nextCounter[1] += count
			} else if shittyLog10(v)%2 == 0 {
				s := fmt.Sprintf("%d", v)
				left, _ := strconv.Atoi(s[0 : len(s)/2])
				right, _ := strconv.Atoi(s[len(s)/2:])
				nextCounter[left] += count
				nextCounter[right] += count
			} else {
				nextCounter[v*2024] += count
			}
		}

		counter, nextCounter = nextCounter, counter
		clear(nextCounter)
	}

	for _, count := range counter {
		p2 += count
	}

	fmt.Printf("%d\n%d\n", p1, p2)
}
