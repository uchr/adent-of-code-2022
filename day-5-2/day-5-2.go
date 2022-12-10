package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func initStacks(scanner *bufio.Scanner) [][]byte {
	scanner.Scan()
	line := scanner.Text()

	stacks := make([][]byte, (len(line)+1)/4)
	parseLine(stacks, line)

	return stacks
}

func parseLine(stacks [][]byte, line string) {
	for i := 1; i < len(line); i += 4 {
		if line[i] != ' ' && !unicode.IsNumber(rune(line[i])) {
			stacks[(i-1)/4] = append(stacks[(i-1)/4], line[i])
		}
	}
}

func parseStacks(scanner *bufio.Scanner) [][]byte {
	stacks := initStacks(scanner)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		parseLine(stacks, line)
	}

	for _, stack := range stacks {
		for i := 0; i < len(stack)/2; i++ {
			stack[i], stack[len(stack)-i-1] = stack[len(stack)-i-1], stack[i]
		}
	}

	return stacks
}

func printStack(stacks [][]byte) {
	for i, stack := range stacks {
		fmt.Printf("%d: ", i)
		for _, c := range stack {
			fmt.Print(string(c))
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func move(stacks [][]byte, num, from, to int) {
	stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-num:]...)
	stacks[from] = stacks[from][:len(stacks[from])-num]
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}

	scanner := bufio.NewScanner(file)
	stacks := parseStacks(scanner)

	printStack(stacks)

	for scanner.Scan() {
		line := scanner.Text()
		var num, from, to int
		_, err := fmt.Fscanf(strings.NewReader(line), "move %d from %d to %d", &num, &from, &to)
		if err != nil {
			log.Panic(err)
		}

		move(stacks, num, from-1, to-1)
		printStack(stacks)
	}

	for _, stack := range stacks {
		fmt.Print(string(stack[len(stack)-1]))
	}
}
