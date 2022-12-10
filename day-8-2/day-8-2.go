package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/gookit/color"
)

type Tree struct {
	xPositive, xNegative, yPositive, yNegative int
}

type Field struct {
	data [][]byte
	tree [][]Tree
	w, h int
}

const (
	PrintField  = 0
	PrintPoints = 1
)

type CellColor struct {
	x, y int
	c    color.Color
}

func (f *Field) Print(param int, colors ...CellColor) {
	for x := range f.tree {
		for y := range f.tree[x] {
			for _, c := range colors {
				if c.x == x && c.y == y {
					color.Set(c.c)
				}
			}

			switch param {
			case PrintField:
				fmt.Printf("%c ", f.data[x][y])
			case PrintPoints:
				fmt.Printf("%d ", f.GetPoints(x, y))
			}
			color.Reset()
		}
		fmt.Println()
	}

	fmt.Println()
}

func (f *Field) PrintCell(param int, x, y int) {
	f.Print(param,
		CellColor{
			x: x,
			y: y,
			c: color.FgGreen,
		},
		CellColor{
			x: f.tree[x][y].xPositive,
			y: y,
			c: color.FgLightRed,
		},
		CellColor{
			x: f.tree[x][y].xNegative,
			y: y,
			c: color.FgLightRed,
		},
		CellColor{
			x: x,
			y: f.tree[x][y].yPositive,
			c: color.FgLightMagenta,
		},
		CellColor{
			x: x,
			y: f.tree[x][y].yNegative,
			c: color.FgLightMagenta,
		},
	)
}

func (f Field) GetPoints(x, y int) int {
	return (x - f.tree[x][y].xPositive) *
		(y - f.tree[x][y].yPositive) *
		(f.tree[x][y].xNegative - x) *
		(f.tree[x][y].yNegative - y)
}

type ClosestTree struct {
	points []int
}

func (t *ClosestTree) GetClosest(height byte) int {
	return t.points[height-byte('0')]
}

func (t *ClosestTree) UpdateClosest(coord int, height byte) {
	for i := 0; i <= int(height)-int('0'); i++ {
		t.points[i] = coord
	}
}

func NewClosestTree(defaultValue int) ClosestTree {
	t := ClosestTree{points: make([]int, 10)}
	for i := range t.points {
		t.points[i] = defaultValue
	}
	return t
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		log.Panic(err)
	}

	data = bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n"))

	f := Field{}
	f.data = bytes.Split(data, []byte("\n"))
	f.data = f.data[:len(f.data)-1] // remove end blank line

	f.w, f.h = len(f.data), len(f.data[0])
	f.tree = make([][]Tree, f.w)
	for x := 0; x < f.w; x++ {
		f.tree[x] = make([]Tree, f.h)
	}

	for x := 0; x < f.w; x++ {
		closestTree := NewClosestTree(0)
		for y := 0; y < f.h; y++ {
			f.tree[x][y].yPositive = closestTree.GetClosest(f.data[x][y])
			closestTree.UpdateClosest(y, f.data[x][y])
		}
	}

	for x := 0; x < f.w; x++ {
		closestTree := NewClosestTree(f.h - 1)
		for y := f.h - 1; y >= 0; y-- {
			f.tree[x][y].yNegative = closestTree.GetClosest(f.data[x][y])
			closestTree.UpdateClosest(y, f.data[x][y])
		}
	}

	for y := 0; y < f.h; y++ {
		closestTree := NewClosestTree(0)
		for x := 0; x < f.w; x++ {
			f.tree[x][y].xPositive = closestTree.GetClosest(f.data[x][y])
			closestTree.UpdateClosest(x, f.data[x][y])
		}
	}

	for y := 0; y < f.h; y++ {
		closestTree := NewClosestTree(f.w - 1)
		for x := f.w - 1; x >= 0; x-- {
			f.tree[x][y].xNegative = closestTree.GetClosest(f.data[x][y])
			closestTree.UpdateClosest(x, f.data[x][y])
		}
	}

	maxPoints := 0
	for x := 1; x < f.w-1; x++ {
		for y := 1; y < f.h-1; y++ {
			points := f.GetPoints(x, y)
			if maxPoints < points {
				maxPoints = points
			}
		}
	}

	fmt.Println(maxPoints)
}
