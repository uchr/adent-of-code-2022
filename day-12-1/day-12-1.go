package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x, y int
}

type Cell struct {
	isVisited bool
	length    int
	prev      Point
}

func Find(field [][]byte, letter byte) (Point, error) {
	for x := range field {
		for y := range field[x] {
			if field[x][y] == letter {
				return Point{x: x, y: y}, nil
			}
		}
	}

	return Point{}, fmt.Errorf("letter '%c' not found", letter)
}

func Adjacent(field [][]byte, p Point) []Point {
	result := make([]Point, 0)
	if p.x > 0 {
		result = append(result, Point{p.x - 1, p.y})
	}
	if p.y > 0 {
		result = append(result, Point{p.x, p.y - 1})
	}
	if p.x < len(field)-1 {
		result = append(result, Point{p.x + 1, p.y})
	}
	if p.y < len(field[0])-1 {
		result = append(result, Point{p.x, p.y + 1})
	}
	return result
}

func Print(fieldData [][]Cell, s, e Point) {
	for x := range fieldData {
		for y := range fieldData[x] {
			if s.x == x && s.y == y {
				fmt.Print("S")
				continue
			}

			if e.x == x && e.y == y {
				fmt.Print("E")
				continue
			}

			if !fieldData[x][y].isVisited {
				fmt.Print(".")
				continue
			}

			if fieldData[x][y].prev.y < y {
				fmt.Print("<")
				continue
			} else if fieldData[x][y].prev.y > y {
				fmt.Print(">")
				continue
			}

			if fieldData[x][y].prev.x < x {
				fmt.Print("^")
				continue
			} else if fieldData[x][y].prev.x > x {
				fmt.Print("V")
				continue

			}
		}
		fmt.Print("\n")
	}

	fmt.Print("\n")
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		log.Panic(err)
	}

	data = bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n"))

	field := bytes.Split(data, []byte("\n"))
	if len(field[len(field)-1]) == 0 {
		field = field[:len(field)-1]
	}
	fmt.Println(field)

	fieldData := make([][]Cell, len(field))
	for x := range fieldData {
		fieldData[x] = make([]Cell, len(field[x]))
	}

	startPoint, err := Find(field, 'S')
	if err != nil {
		log.Panic(err)
	}

	endPoint, err := Find(field, 'E')
	if err != nil {
		log.Panic(err)
	}

	field[startPoint.x][startPoint.y] = 'a' - 1
	field[endPoint.x][endPoint.y] = 'z' + 1

	wave := []Point{startPoint}
	fieldData[startPoint.x][startPoint.y].length = 0
	fieldData[startPoint.x][startPoint.y].isVisited = true

	waveIndex := 1
	Print(fieldData, startPoint, endPoint)
	for len(wave) > 0 {
		nextWave := make([]Point, 0)
		for _, p := range wave {
			for _, ap := range Adjacent(field, p) {
				if fieldData[ap.x][ap.y].isVisited {
					continue
				}

				if field[ap.x][ap.y] > field[p.x][p.y]+1 {
					continue
				}

				nextWave = append(nextWave, ap)
				fieldData[ap.x][ap.y].prev = p
				fieldData[ap.x][ap.y].length = waveIndex
				fieldData[ap.x][ap.y].isVisited = true
			}
		}

		waveIndex++
		wave = nextWave
		fmt.Println("Wave:", waveIndex)
		Print(fieldData, startPoint, endPoint)
	}

	fmt.Println(fieldData[endPoint.x][endPoint.y].length)
}
