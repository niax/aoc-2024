package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/niax/aoc-2024/internal/collections"
)

type Grid = collections.SliceGrid[int]
type Point = collections.Point2D

type pointDistance struct {
	p Point
	d int
}

func pathCountToPoint(grid *Grid, startPoint Point) *collections.SliceGrid[int] {
	reachable := collections.NewSliceGrid[int](grid.Width(), grid.Height())

	frontier := make([]Point, 1, grid.Width())
	frontier[0] = startPoint

	for len(frontier) > 0 {
		pd := frontier[0]
		cur := reachable.AtPoint(pd)
		reachable.Set(pd.X, pd.Y, *cur+1)
		wantedHeight := *grid.AtPoint(pd) + 1
		frontier = frontier[1:]

		for _, dir := range collections.Point2D_CardinalDirections {
			nextPoint := pd.Add(dir)
			nextGridVal := grid.AtPoint(nextPoint)
			if nextGridVal == nil {
				continue
			}

			if *nextGridVal == wantedHeight {
				frontier = append(frontier, nextPoint)
			}
		}
	}

	return reachable
}

func main() {
	inputFd, err := os.Open("inputs/10")
	if err != nil {
		panic(err)
	}
	defer inputFd.Close()

	input := bufio.NewScanner(inputFd)

	startPoints := make([]collections.Point2D, 0, 256)
	endPoints := make([]collections.Point2D, 0, 256)
	grid := collections.NewSliceGrid[int](58, 58)
	y := 0
	for input.Scan() {
		line := input.Text()

		for x := range line {
			cell, _ := strconv.Atoi(line[x : x+1])
			grid.Set(x, y, cell)
			if cell == 0 {
				startPoints = append(startPoints, collections.Point2D{
					X: x,
					Y: y,
				})
			} else if cell == 9 {
				endPoints = append(endPoints, collections.Point2D{
					X: x,
					Y: y,
				})
			}
		}

		y++
	}

	p1 := 0
	p2 := 0
	for i := range startPoints {
		distances := pathCountToPoint(grid, startPoints[i])
		for j := range endPoints {
			c := *distances.AtPoint(endPoints[j])
			if c != 0 {
				p1++
				p2 += *distances.AtPoint(endPoints[j])
			}
		}
	}

	fmt.Printf("%d\n%d\n", p1, p2)
}
