package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
)

type charGrid [][]byte

func (g charGrid) Print() {
	for y := range g {
		for x := range g[y] {
			fmt.Print(string(rune(g[y][x])))
		}
		fmt.Print("\n")
	}
}

func (g charGrid) At(x, y int) (byte, bool) {
	if y < 0 || y >= len(g) {
		return 0, false
	}

	row := g[y]
	if x < 0 || x >= len(row) {
		return 0, false
	}

	return row[x], true
}

func (g charGrid) IterDir(x, y, dx, dy int) iter.Seq[byte] {
	return func(yield func(byte) bool) {
		for {
			b, ok := g.At(x, y)
			if !ok {
				return
			}
			if !yield(b) {
				return
			}
			x += dx
			y += dy
		}
	}
}

func main() {
	inputFd, err := os.Open("inputs/04")
	if err != nil {
		panic(err)
	}
	defer inputFd.Close()

	input := bufio.NewScanner(inputFd)
	grid := make(charGrid, 0, 1024)
	for input.Scan() {
		line := input.Text()
		row := make([]byte, len(line))
		for i := range line {
			row[i] = line[i]
		}
		grid = append(grid, row)
	}

	p1 := 0
	p2 := 0
	acc := make([]byte, 4)
	for y := range grid {
		for x := range grid[y] {
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dx == 0 && dy == 0 {
						continue
					}
					acc = acc[:0]
					for c := range grid.IterDir(x, y, dx, dy) {
						acc = append(acc, c)
						if len(acc) == 4 {
							if acc[0] == 88 && acc[1] == 77 && acc[2] == 65 && acc[3] == 83 {
								p1++
							}
							break
						}
					}
				}
			}

			if grid[y][x] == 65 {
				topLeft, ok := grid.At(x-1, y-1)
				if !ok {
					goto exit
				}
				topRight, ok := grid.At(x+1, y-1)
				if !ok {
					goto exit
				}
				bottomLeft, ok := grid.At(x-1, y+1)
				if !ok {
					goto exit
				}
				bottomRight, ok := grid.At(x+1, y+1)
				if !ok {
					goto exit
				}
				hasTopLeftMas := (topLeft == 77 && bottomRight == 83) || (topLeft == 83 && bottomRight == 77)
				hasTopRightMas := (topRight == 77 && bottomLeft == 83) || (topRight == 83 && bottomLeft == 77)
				if hasTopLeftMas && hasTopRightMas {
					p2++
				}
			}
		exit:
		}
	}

	fmt.Printf("%d\n%d\n", p1, p2)

}
