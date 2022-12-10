package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type result int

const (
	win  result = 6
	draw result = 3
	lose result = 0
)

type handScore int

const (
	rock     handScore = 1
	paper    handScore = 2
	scissors handScore = 3
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}

	// A/X - Rock
	// B/Y - Paper
	// C/Z - Scissors
	cases := map[string]int{
		"A X": int(draw) + int(rock),
		"A Y": int(win) + int(paper),
		"A Z": int(lose) + int(scissors),
		"B X": int(lose) + int(rock),
		"B Y": int(draw) + int(paper),
		"B Z": int(win) + int(scissors),
		"C X": int(win) + int(rock),
		"C Y": int(lose) + int(paper),
		"C Z": int(draw) + int(scissors),
	}

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		score += int(cases[line])
	}

	fmt.Println(score)
}
