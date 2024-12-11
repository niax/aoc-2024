package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/knz/go-ilog10"
)

func shittyLog10(i int) int {
	return int(ilog10.FastUint64Log10(uint64(i)))
}

var pow10Lookup = []int{
	1, 10, 100, 1000, 10_000, 100_000, 1_000_000, 10_000_000,
}

func shittyPow10(i int) int {
	return pow10Lookup[i]
}

func main() {
	inputFd, err := os.Open("inputs/11")
	if err != nil {
		panic(err)
	}
	defer inputFd.Close()

	counter := make(map[int]int, 4096)
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

	nextCounter := make(map[int]int, 4096)
	for step := 0; step < 75; step++ {
		if step == 25 {
			for _, count := range counter {
				p1 += count
			}
		}

		for v, count := range counter {
			log10 := shittyLog10(v)
			if v == 0 {
				nextCounter[1] += count
			} else if log10%2 == 0 {
				mask := shittyPow10(log10 / 2)
				left := v / mask
				right := v % mask
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
