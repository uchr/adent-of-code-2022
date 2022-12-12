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

type ArgType byte

const (
	VariableArgType ArgType = iota
	ConstantArgType
)

type Operation struct {
	LeftType  ArgType
	LeftValue int64

	RightType  ArgType
	RightValue int64

	Sign byte
}

type Monkey struct {
	Items     []int64
	Operation Operation

	Divider     int64
	TrueResult  int
	FalseResult int

	InspectationCount int64
}

func (m *Monkey) Evaluate(variable int64) int64 {
	var left, right int64

	switch m.Operation.LeftType {
	case ConstantArgType:
		left = m.Operation.LeftValue
	case VariableArgType:
		left = variable
	}

	switch m.Operation.RightType {
	case ConstantArgType:
		right = m.Operation.RightValue
	case VariableArgType:
		right = variable
	}

	switch m.Operation.Sign {
	case '+':
		return left + right
	case '-':
		return left - right
	case '*':
		return left * right
	case '/':
		return left / right
	default:
		log.Panic("unknown operation")
	}

	return 0
}

func (m *Monkey) Test(value int64) int {
	if value%m.Divider == 0 {
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
				n, err := strconv.ParseInt(number, 10, 0)
				if err != nil {
					log.Panic(err)
				}
				currentMonkey.Items = append(currentMonkey.Items, n)
			}
		}

		if strings.Contains(line, "Operation:") {
			exprs := strings.Split(line[strings.Index(line, ":")+2:], " ")

			if exprs[2] != "old" {
				value, err := strconv.ParseInt(exprs[2], 10, 0)
				if err != nil {
					log.Panic(err)
				}
				currentMonkey.Operation.LeftType = ConstantArgType
				currentMonkey.Operation.LeftValue = value
			} else {
				currentMonkey.Operation.LeftType = VariableArgType
			}

			if exprs[4] != "old" {
				value, err := strconv.ParseInt(exprs[4], 10, 0)
				if err != nil {
					log.Panic(err)
				}
				currentMonkey.Operation.RightType = ConstantArgType
				currentMonkey.Operation.RightValue = value
			} else {
				currentMonkey.Operation.RightType = VariableArgType
			}

			currentMonkey.Operation.Sign = exprs[3][0]
		}

		if strings.Contains(line, "Test:") {
			number := line[strings.LastIndex(line, " ")+1:]
			n, err := strconv.ParseInt(number, 10, 0)
			if err != nil {
				log.Panic(err)
			}
			currentMonkey.Divider = n
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

	dividers := int64(1)
	for _, m := range monkeys {
		dividers *= m.Divider
	}

	for round := 1; round <= 10000; round++ {
		for monkeyIndex := range monkeys {
			monkey := &monkeys[monkeyIndex]
			items := monkey.Items
			monkey.Items = make([]int64, 0)
			for _, item := range items {
				monkey.InspectationCount++
				worryLevel := monkey.Evaluate(item)
				worryLevel %= dividers
				nextMonkey := monkey.Test(worryLevel)
				monkeys[nextMonkey].Items = append(monkeys[nextMonkey].Items, worryLevel)
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].InspectationCount > monkeys[j].InspectationCount })

	fmt.Println(monkeys[0].InspectationCount * monkeys[1].InspectationCount)
}
