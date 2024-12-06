package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"

	"github.com/bits-and-blooms/bitset"
	"github.com/niax/aoc-2024/internal/collections"
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

type direction int

const (
	directionNorth direction = iota
	directionEast
	directionSouth
	directionWest
)

func (d direction) TurnRight() direction {
	n := d + 1
	if n > directionWest {
		n = directionNorth
	}
	return n
}

func (d direction) Delta() (int, int) {
	switch d {
	case directionNorth:
		return 0, -1
	case directionEast:
		return 1, 0
	case directionSouth:
		return 0, 1
	case directionWest:
		return -1, 0
	default:
		panic("NO!")
	}
}

type solver struct {
	grid        *SliceGrid[bool]
	guardStartX int
	guardStartY int

	visited    collections.Set[int]
	visitedDir bitset.BitSet
}

func NewSolver(grid *SliceGrid[bool], guardStartX int, guardStartY int) *solver {
	visited := collections.NewSetWithCapacity[int](1024 * 10)
	visitedDir := bitset.New(64 * 1024)
	return &solver{
		grid:        grid,
		guardStartX: guardStartX,
		guardStartY: guardStartY,

		visited:    visited,
		visitedDir: *visitedDir,
	}
}

func (s *solver) canEscape() bool {
	shouldUpdateVisisted := len(s.visited) == 0
	s.visitedDir.ClearAll()

	posX := s.guardStartX
	posY := s.guardStartY
	dir := directionNorth
	dx, dy := dir.Delta()
	for posX >= 0 && posX < s.grid.width && posY >= 0 && posY < s.grid.height {
		if shouldUpdateVisisted {
			s.visited.Add((posX << 16) + posY)
		}
		visitedDirEnc := uint((int(dir) << 16) | (posX << 8) | posY)
		if s.visitedDir.Test(visitedDirEnc) {
			return false
		}
		s.visitedDir.Set(visitedDirEnc)

		nextX := posX + dx
		nextY := posY + dy
		nextSpot := s.grid.At(nextX, nextY)
		if nextSpot != nil && *nextSpot {
			// Rotate
			dir = dir.TurnRight()
			dx, dy = dir.Delta()
		} else {
			posX = nextX
			posY = nextY
		}
	}

	return true
}

func main() {
	inputFd, err := os.Open("inputs/06")
	if err != nil {
		panic(err)
	}
	defer inputFd.Close()

	grid := NewSliceGrid[bool](130, 130)
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

	slvr := NewSolver(grid, posX, posY)
	slvr.canEscape()
	p1 := len(slvr.visited)
	pathPlaces := maps.Clone(slvr.visited)

	p2 := 0

	for encoded := range pathPlaces {
		x := encoded >> 16
		y := encoded & 0xffff
		slvr.grid.Set(x, y, true)
		if !slvr.canEscape() {
			p2++
		}
		slvr.grid.Set(x, y, false)
	}

	fmt.Printf("%d\n%d\n", p1, p2)
}
