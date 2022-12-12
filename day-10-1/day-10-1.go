package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Runtime struct {
	value int
	cycle int

	result int
}

func NewRuntime() Runtime {
	return Runtime{
		value: 1,
		cycle: 1,
	}
}

func (r *Runtime) IncCycle() {
	if r.cycle%20 == 0 && (r.cycle == 20 || ((r.cycle-20)%40 == 0)) {
		r.result += r.cycle * r.value
		fmt.Println(r.value)
	}
	r.cycle++
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

	fmt.Println(r.result)
}
