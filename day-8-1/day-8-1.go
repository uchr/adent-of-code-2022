package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/gookit/color"
)

type Field struct {
	data       [][]byte
	visibility [][]bool
	w, h       int
}

func (f *Field) Print(cx, cy int) {
	for x := range f.visibility {
		for y := range f.visibility[x] {
			if cx == x && cy == y {
				color.Set(color.FgGreen)
			} else if f.visibility[x][y] {
				color.Set(color.FgRed)
			}
			fmt.Printf("%c", f.data[x][y])
			color.Reset()
		}
		fmt.Println()
	}
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
	f.visibility = make([][]bool, f.w)
	for x := 0; x < f.w; x++ {
		f.visibility[x] = make([]bool, f.h)
	}

	count := f.w*2 + f.h*2 - 4 // corner trees
	fmt.Println(count)
	for x := 1; x < f.w-1; x++ {
		curHeight := f.data[x][0]
		for y := 1; y <= f.h-2; y++ {
			if curHeight < f.data[x][y] {
				if !f.visibility[x][y] {
					f.visibility[x][y] = true
					count++
				}
				curHeight = f.data[x][y]
			}
		}
	}

	for x := 1; x < f.w-1; x++ {
		curHeight := f.data[x][f.h-1]
		for y := f.h - 2; y >= 1; y-- {
			if curHeight < f.data[x][y] {
				if !f.visibility[x][y] {
					f.visibility[x][y] = true
					count++
				}
				curHeight = f.data[x][y]
			}
		}
	}

	for y := 1; y < f.h-1; y++ {
		curHeight := f.data[0][y]
		for x := 1; x <= f.w-2; x++ {
			if curHeight < f.data[x][y] {
				if !f.visibility[x][y] {
					f.visibility[x][y] = true
					count++
				}
				curHeight = f.data[x][y]
			}
		}
	}

	for y := 1; y < f.h-1; y++ {
		curHeight := f.data[f.w-1][y]
		for x := f.w - 2; x >= 1; x-- {
			if curHeight < f.data[x][y] {
				if !f.visibility[x][y] {
					f.visibility[x][y] = true
					count++
				}
				curHeight = f.data[x][y]
			}
		}
	}

	f.Print(-1, -1)

	fmt.Println(count)
}
