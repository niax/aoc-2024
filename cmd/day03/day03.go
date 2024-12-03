package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type commandType int

const (
	commandTypeUnknown commandType = iota
	commandTypeMul
	commandTypeDo
	commandTypeDont
)

type command struct {
	cmd    commandType
	param1 int
	param2 int
}

type parseState int

const (
	parseStateInitial parseState = iota
	parseStateSeenMUL_PAREN
	parseStateSeenMUL_COMMA
)

func parse(s string) []command {
	commands := make([]command, 0, 128)
	state := parseStateInitial
	accum1 := make([]rune, 0, len(s))
	accum2 := make([]rune, 0, len(s))
	i := 0
	for i < len(s) {
		c := rune(s[i])
		switch state {
		case parseStateInitial:
			x := s[i:]
			if strings.HasPrefix(x, "mul(") {
				state = parseStateSeenMUL_PAREN
				i += 4
			} else if strings.HasPrefix(x, "do()") {
				commands = append(commands, command{
					cmd: commandTypeDo,
				})
				i += 4
			} else if strings.HasPrefix(x, "don't()") {
				commands = append(commands, command{
					cmd: commandTypeDont,
				})
				i += 7
			} else {
				i += 1
			}
		case parseStateSeenMUL_PAREN:
			if unicode.IsDigit(c) {
				accum1 = append(accum1, c)
			} else if c == ',' {
				state = parseStateSeenMUL_COMMA
			} else {
				accum1 = accum1[:0]
				accum2 = accum2[:0]
				state = parseStateInitial
			}
			i += 1
		case parseStateSeenMUL_COMMA:
			if unicode.IsDigit(c) {
				accum2 = append(accum2, c)
			} else if c == ')' {
				p1, _ := strconv.Atoi(string(accum1))
				p2, _ := strconv.Atoi(string(accum2))
				commands = append(commands, command{
					cmd:    commandTypeMul,
					param1: p1,
					param2: p2,
				})
				state = parseStateInitial
				accum1 = accum1[:0]
				accum2 = accum2[:0]
			} else {
				accum1 = accum1[:0]
				accum2 = accum2[:0]
				state = parseStateInitial
			}
			i += 1
		default:
			panic("unknown parse state")
		}
	}

	return commands
}

func main() {
	inputFd, err := os.Open("inputs/03")
	if err != nil {
		panic(err)
	}
	defer inputFd.Close()

	p1 := 0
	p2 := 0
	input := bufio.NewScanner(inputFd)
	enable := true
	for input.Scan() {
		line := input.Text()
		commands := parse(line)
		for i := range commands {
			switch commands[i].cmd {
			case commandTypeMul:
				mul := commands[i].param1 * commands[i].param2
				p1 += mul
				if enable {
					p2 += mul
				}
			case commandTypeDo:
				enable = true
			case commandTypeDont:
				enable = false
			}
		}
	}

	fmt.Printf("%d\n%d\n", p1, p2)
}
