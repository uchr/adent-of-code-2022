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

	// A - Rock
	// B - Paper
	// C - Scissors
	// X - Lose
	// Y - Draw
	// Z - Win
	cases := map[string]int{
		"A X": int(lose) + int(scissors),
		"A Y": int(draw) + int(rock),
		"A Z": int(win) + int(paper),
		"B X": int(lose) + int(rock),
		"B Y": int(draw) + int(paper),
		"B Z": int(win) + int(scissors),
		"C X": int(lose) + int(paper),
		"C Y": int(draw) + int(scissors),
		"C Z": int(win) + int(rock),
	}

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		score += cases[line]
	}

	fmt.Println(score)
}
