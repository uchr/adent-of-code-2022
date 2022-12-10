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

type Rope struct {
	head, tail Point
}

func Abs(value int) int {
	if value >= 0 {
		return value
	}

	return -value
}

func (r *Rope) IsTouched() bool {
	return Abs(r.head.x-r.tail.x) <= 1 && Abs(r.head.y-r.tail.y) <= 1
}

func (r *Rope) UpdateTail() {
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

func (r *Rope) MoveHead(command string) {
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

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}

	positions := make(map[Point]bool)

	r := Rope{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		command := line[0]
		steps, err := strconv.Atoi(line[1])
		if err != nil {
			log.Panic(err)
		}

		for i := 0; i < steps; i++ {
			r.MoveHead(command)
			positions[r.tail] = true
		}
	}

	fmt.Println(len(positions))
}
