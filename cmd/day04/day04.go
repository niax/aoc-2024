package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
)

type SliceGrid[T any] struct {
	data   []T
	width  int
	height int
}

func NewSliceGrid[T any](width, height int) *SliceGrid[T] {
	data := make([]T, width*height)
	return &SliceGrid[T]{
		data:   data,
		width:  width,
		height: height,
	}
}

func (g *SliceGrid[T]) coordIdx(x, y int) int {
	return y*g.width + x
}

func (g *SliceGrid[T]) Width() int {
	return g.width
}

func (g *SliceGrid[T]) Height() int {
	return g.height
}

func (g *SliceGrid[T]) Set(x, y int, v T) {
	if x < 0 || y < 0 {
		panic("OOB!")
	}
	g.data[g.coordIdx(x, y)] = v
}

func (g *SliceGrid[T]) At(x, y int) *T {
	if x < 0 || y < 0 || x >= int(g.width) || y >= int(g.height) {
		return nil
	}
	return &g.data[g.coordIdx(x, y)]
}

func (g *SliceGrid[T]) IterDir(x, y, dx, dy int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			v := g.At(x, y)
			if v == nil {
				return
			}
			if !yield(*v) {
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
	grid := NewSliceGrid[byte](140, 140)
	y := 0
	for input.Scan() {
		line := input.Text()
		for x := range line {
			grid.Set(x, y, line[x])
		}
		y++
	}

	p1 := 0
	p2 := 0
	for y := range grid.Height() {
		for x := range grid.Width() {
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dx == 0 && dy == 0 {
						continue
					}
					i := 0
					for c := range grid.IterDir(x, y, dx, dy) {
						i++
						if i == 1 && c != 88 {
							break
						} else if i == 2 && c != 77 {
							break
						} else if i == 3 && c != 65 {
							break
						} else if i == 4 && c != 83 {
							break
						} else if i > 4 {
							break
						} else if i == 4 && c == 83 {
							p1++
						}
					}
				}
			}

			if *grid.At(x, y) == 65 && x > 0 && x < (grid.Width()-1) && y > 0 && y < (grid.Height()-1) {
				topLeft := *grid.At(x-1, y-1)
				topRight := *grid.At(x+1, y-1)
				bottomLeft := *grid.At(x-1, y+1)
				bottomRight := *grid.At(x+1, y+1)
				hasTopLeftMas := (topLeft == 77 && bottomRight == 83) || (topLeft == 83 && bottomRight == 77)
				hasTopRightMas := (topRight == 77 && bottomLeft == 83) || (topRight == 83 && bottomLeft == 77)
				if hasTopLeftMas && hasTopRightMas {
					p2++
				}
			}
		}
	}

	fmt.Printf("%d\n%d\n", p1, p2)

}
