package collections

import "fmt"

type SliceGrid[T any] struct {
	data   []T
	width  int
	height int
}

type SliceGridIndex struct {
	idx int
	oob bool
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

func (g *SliceGrid[T]) AtPoint(p Point2D) *T {
	return g.At(p.X, p.Y)
}

func (g *SliceGrid[T]) IndexFor(x, y int) SliceGridIndex {
	if x < 0 || y < 0 || x >= int(g.width) || y >= int(g.height) {
		return SliceGridIndex{
			oob: true,
		}
	}
	return SliceGridIndex{
		idx: g.coordIdx(x, y),
	}
}

func (g *SliceGrid[T]) IndexForPoint(p Point2D) SliceGridIndex {
	return g.IndexFor(p.X, p.Y)
}

func (g *SliceGrid[T]) AtIndex(i SliceGridIndex) *T {
	if i.oob {
		return nil
	}
	return &g.data[i.idx]
}

func (g *SliceGrid[T]) SetIndex(i SliceGridIndex, v T) {
	if i.oob {
		panic("OOB!")
	}
	g.data[i.idx] = v
}

func (g *SliceGrid[T]) Print() {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			fmt.Printf("%v", *g.At(x, y))
		}
		fmt.Printf("\n")
	}
}
