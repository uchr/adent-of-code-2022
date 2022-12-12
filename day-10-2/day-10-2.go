package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Abs(value int) int {
	if value >= 0 {
		return value
	}

	return -value
}

type Runtime struct {
	value int
	cycle int

	screen []byte
}

func NewRuntime() Runtime {
	return Runtime{
		value:  1,
		cycle:  0,
		screen: make([]byte, 250),
	}
}

func (r *Runtime) IncCycle() {
	if Abs(r.value-(r.cycle%40)) <= 1 {
		r.screen[r.cycle] = '#'
	} else {
		r.screen[r.cycle] = '.'
	}
	r.cycle++
}

func (r *Runtime) Print() {
	for line := 0; line < 6; line++ {
		fmt.Println(string(r.screen[line*40 : (line+1)*40]))
	}
}

func (r *Runtime) Execute(command []string) {
	switch command[0] {
	case "noop":
		r.IncCycle()
	case "addx":
		r.IncCycle()
		r.IncCycle()
		v, err := strconv.Atoi(command[1])
		if err != nil {
			log.Panic(err)
		}
		r.value += v
	}
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}

	r := NewRuntime()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		r.Execute(command)
	}

	r.Print()
}
