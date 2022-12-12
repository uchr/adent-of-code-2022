package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items     []int
	Operation string

	DividerCondition int
	TrueResult       int
	FalseResult      int

	InspectationCount int
}

func (m Monkey) Evaluate(variable int) int {
	exprs := strings.Split(m.Operation, " ")

	var err error
	var left, right int
	if exprs[2] != "old" {
		left, err = strconv.Atoi(exprs[2])
		if err != nil {
			log.Panic(err)
		}
	} else {
		left = variable
	}

	if exprs[4] != "old" {
		right, err = strconv.Atoi(exprs[4])
		if err != nil {
			log.Panic(err)
		}
	} else {
		right = variable
	}

	switch exprs[3] {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}

	log.Panic("unknown operation")
	return 0
}

func (m Monkey) Test(value int) int {
	if value%m.DividerCondition == 0 {
		return m.TrueResult
	} else {
		return m.FalseResult
	}
}

func (m Monkey) String() string {
	return fmt.Sprintf("%v", m.Items)

}

func ParseMonkeys(scanner *bufio.Scanner) []Monkey {
	result := make([]Monkey, 0)

	currentMonkey := Monkey{}
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Monkey") {
			continue
		}

		if strings.Contains(line, "Starting items:") {
			numbers := strings.Split(line[strings.Index(line, ":")+2:], ", ")
			for _, number := range numbers {
				n, err := strconv.Atoi(number)
				if err != nil {
					log.Panic(err)
				}
				currentMonkey.Items = append(currentMonkey.Items, n)
			}
		}

		if strings.Contains(line, "Operation:") {
			currentMonkey.Operation = line[strings.Index(line, ":")+2:]
		}

		if strings.Contains(line, "Test:") {
			number := line[strings.LastIndex(line, " ")+1:]
			n, err := strconv.Atoi(number)
			if err != nil {
				log.Panic(err)
			}
			currentMonkey.DividerCondition = n
		}

		if strings.Contains(line, "If true:") {
			number := line[strings.LastIndex(line, " ")+1:]
			n, err := strconv.Atoi(number)
			if err != nil {
				log.Panic(err)
			}
			currentMonkey.TrueResult = n
		}

		if strings.Contains(line, "If false:") {
			number := line[strings.LastIndex(line, " ")+1:]
			n, err := strconv.Atoi(number)
			if err != nil {
				log.Panic(err)
			}
			currentMonkey.FalseResult = n
		}

		if line == "" {
			result = append(result, currentMonkey)
			currentMonkey = Monkey{}
		}
	}
	result = append(result, currentMonkey)

	return result
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}

	scanner := bufio.NewScanner(file)

	monkeys := ParseMonkeys(scanner)

	for round := 1; round <= 20; round++ {
		for monkeyIndex := range monkeys {
			monkey := &monkeys[monkeyIndex]
			items := monkey.Items
			monkey.Items = make([]int, 0)
			for _, item := range items {
				monkey.InspectationCount++
				worryLevel := monkey.Evaluate(item)
				worryLevel /= 3
				nextMonkey := monkey.Test(worryLevel)
				monkeys[nextMonkey].Items = append(monkeys[nextMonkey].Items, worryLevel)
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].InspectationCount > monkeys[j].InspectationCount })

	fmt.Println(monkeys[0].InspectationCount * monkeys[1].InspectationCount)

}
