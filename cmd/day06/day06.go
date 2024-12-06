package main

import (
	"bufio"
	"fmt"
	"github.com/niax/aoc-2024/internal/collections"
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

func canEscape(grid *SliceGrid[bool], posX, posY int) *collections.Set[int] {
	dy := -1
	dx := 0
	visited := collections.NewSet[int]()
	visitedDir := collections.NewSet[int]()
	for posX >= 0 && posX < grid.Width() && posY >= 0 && posY < grid.Height() {
		visited.Add((posX << 16) | posY)
		dirEnc := 0
		if dx == 0 && dy == -1 {
			dirEnc = 1
		} else if dx == 1 && dy == 0 {
			dirEnc = 2
		} else if dx == 0 && dy == 1 {
			dirEnc = 3
		} else if dx == -1 && dy == 0 {
			dirEnc = 4
		} else {
			panic("NOO!")
		}
		visitedDirEnc := (dirEnc << 32) | (posX << 16) | posY
		if visitedDir.Contains(visitedDirEnc) {
			return nil
		}
		visitedDir.Add(visitedDirEnc)

		nextX := posX + dx
		nextY := posY + dy
		nextSpot := grid.At(nextX, nextY)
		if nextSpot != nil && *nextSpot {
			// Rotate
			if dx == 0 && dy == -1 {
				dx = 1
				dy = 0
			} else if dx == 1 && dy == 0 {
				dx = 0
				dy = 1
			} else if dx == 0 && dy == 1 {
				dx = -1
				dy = 0
			} else if dx == -1 && dy == 0 {
				dx = 0
				dy = -1
			} else {
				panic("NOO!")
			}
		} else {
			posX = nextX
			posY = nextY
		}
	}

	return &visited
}

func main() {
	inputFd, err := os.Open("inputs/06")
	if err != nil {
		panic(err)
	}
	defer inputFd.Close()

	grid := NewSliceGrid[bool](130, 130)
	//grid := NewSliceGrid[bool](10, 10)
	input := bufio.NewScanner(inputFd)
	posX := 0
	posY := 0
	y := 0
	for input.Scan() {
		line := input.Text()
		for x := range line {
			if line[x] == '#' {
				grid.Set(x, y, true)
			}
			if line[x] == '^' {
				posX = x
				posY = y
			}
		}
		y++
	}

	pathPlaces := *canEscape(grid, posX, posY)
	p1 := len(pathPlaces)
	p2 := 0

	for encoded := range pathPlaces {
		x := encoded >> 16
		y := encoded & 0xffff
		grid.Set(x, y, true)
		if canEscape(grid, posX, posY) == nil {
			p2++
		}
		grid.Set(x, y, false)
	}

	fmt.Printf("%d\n%d\n", p1, p2)
}
