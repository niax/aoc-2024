package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	cmd commandType
	param1 int
	param2 int
}

type parseState int
const (
	parseStateInitial parseState = iota
	parseStateSeenM 
	parseStateSeenMU
	parseStateSeenMUL
	parseStateSeenMUL_PAREN
	parseStateSeenMUL_COMMA
	parseStateSeenD
	parseStateSeenDO
	parseStateSeenDO_PAREN
	parseStateSeenDON
	parseStateSeenDONX
	parseStateSeenDONXT
	parseStateSeenDONXT_PAREN
)

func parse(s string) []command {
	commands := make([]command, 0, 128)
	state := parseStateInitial
	accum1 := make([]rune, 0, len(s))
	accum2 := make([]rune, 0, len(s))
	for _, c := range s {
		switch state {
		case parseStateInitial:
			switch c {
			case 'm':
				state = parseStateSeenM
			case 'd':
				state = parseStateSeenD
			}
		case parseStateSeenM:
			switch c {
			case 'u':
				state = parseStateSeenMU
			default:
				state = parseStateInitial
			}
		case parseStateSeenMU:
			switch c {
			case 'l':
				state = parseStateSeenMUL
			default:
				state = parseStateInitial
			}
		case parseStateSeenMUL:
			switch c {
			case '(':
				state = parseStateSeenMUL_PAREN
			default:
				state = parseStateInitial
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
		case parseStateSeenMUL_COMMA:
			if unicode.IsDigit(c) {
				accum2 = append(accum2, c)
			} else if c == ')' {
				p1, _ := strconv.Atoi(string(accum1))
				p2, _ := strconv.Atoi(string(accum2))
				commands = append(commands, command{
					cmd: commandTypeMul,
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
		case parseStateSeenD:
			switch c {
			case 'o':
				state = parseStateSeenDO
			default:
				state = parseStateInitial
			}
		case parseStateSeenDO:
			switch c {
			case '(':
				state = parseStateSeenDO_PAREN
			case 'n':
				state = parseStateSeenDON
			default:
				state = parseStateInitial
			}
		case parseStateSeenDO_PAREN:
			if c == ')' {
				commands = append(commands, command{
					cmd: commandTypeDo,
				})
			}
			state = parseStateInitial
		case parseStateSeenDON:
			switch c {
			case '\'':
				state = parseStateSeenDONX
			default:
				state = parseStateInitial
			}
		case parseStateSeenDONX:
			switch c {
			case 't':
				state = parseStateSeenDONXT
			default:
				state = parseStateInitial
			}
		case parseStateSeenDONXT:
			switch c {
			case '(':
				state = parseStateSeenDONXT_PAREN
			default:
				state = parseStateInitial
			}
		case parseStateSeenDONXT_PAREN:
			if c == ')' {
				commands = append(commands, command{
					cmd: commandTypeDont,
				})
			}
			state = parseStateInitial
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
