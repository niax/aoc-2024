package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputFd, err := os.Open("inputs/03")
	if err != nil {
		panic(err)
	}
	defer inputFd.Close()

	re := regexp.MustCompile("(mul\\((\\d+),(\\d+)\\)|do\\(\\)|don't\\(\\))")
	re.Match([]byte("x"))

	p1 := 0
	p2 := 0
	input := bufio.NewScanner(inputFd)
	enable := true
	for input.Scan() {
		line := input.Text()
		submatches := re.FindAllStringSubmatch(line, -1)
		for i := range submatches {
			submatch := submatches[i]
			if submatch[0] == "do()" {
				enable = true
			} else if submatch[0] == "don't()" {
				enable = false
			}
			a, _ := strconv.Atoi(submatch[2])
			b, _ := strconv.Atoi(submatch[3])
			p1 += (a * b)
			if enable {
				p2 += (a * b)
			}
		}
	}

	fmt.Printf("%d\n%d\n", p1, p2)
}
