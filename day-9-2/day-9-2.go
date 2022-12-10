package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type RopePart struct {
	head, tail Point
}

func Abs(value int) int {
	if value >= 0 {
		return value
	}

	return -value
}

func (r *RopePart) IsTouched() bool {
	return Abs(r.head.x-r.tail.x) <= 1 && Abs(r.head.y-r.tail.y) <= 1
}

func (r *RopePart) UpdateTail() {
	if r.IsTouched() {
		return
	}

	if r.head.y < r.tail.y {
		r.tail.y -= 1
	} else if r.head.y > r.tail.y {
		r.tail.y += 1
	}

	if r.head.x < r.tail.x {
		r.tail.x -= 1
	} else if r.head.x > r.tail.x {
		r.tail.x += 1
	}
}

func (r *RopePart) MoveHead(command string) {
	switch command {
	case "R":
		r.head.x += 1
	case "L":
		r.head.x -= 1
	case "U":
		r.head.y += 1
	case "D":
		r.head.y -= 1
	}

	r.UpdateTail()
}

func (r *RopePart) SetHead(p Point) {
	r.head = p
	r.UpdateTail()
}

type Rope struct {
	parts []RopePart
}

func NewRope(partCount int) Rope {
	return Rope{parts: make([]RopePart, partCount)}
}

func (r *Rope) MoveHead(command string) {
	r.parts[0].MoveHead(command)

	for i := 1; i < len(r.parts); i++ {
		r.parts[i].SetHead(r.parts[i-1].tail)
	}
}

func (r *Rope) TailPosition() Point {
	return r.parts[len(r.parts)-1].tail
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}

	positions := make(map[Point]bool)

	step := 0
	r := NewRope(9)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		command := line[0]
		steps, err := strconv.Atoi(line[1])
		if err != nil {
			log.Panic(err)
		}

		for i := 0; i < steps; i++ {
			step++
			r.MoveHead(command)
			positions[r.TailPosition()] = true
		}
	}

	fmt.Println(len(positions))
}
