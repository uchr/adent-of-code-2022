package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
)

func getLine(line string) []byte {
	result := []byte(line)
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
}

func findCommonLetter(group [][]byte) (byte, error) {
	inds := make([]int, len(group))

	for {
		// loop cond
		for j := range inds {
			if inds[j] >= len(group[j]) {
				return 0, errors.New("reached end of the line")
			}
		}

		// body
		isSameLetter := true
		for j := range inds {
			if group[j][inds[j]] != group[(j+1)%len(group)][inds[(j+1)%len(group)]] {
				isSameLetter = false
				break
			}
		}
		if isSameLetter {
			return group[0][inds[0]], nil
		}

		// loop increment
		lowerInd := 0
		for j := range inds {
			if group[j][inds[j]] < group[lowerInd][inds[lowerInd]] {
				lowerInd = j
			}
		}
		inds[lowerInd]++
	}

	return 0, errors.New("solution not found")
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}

	sum := 0
	scanner := bufio.NewScanner(file)

	const maxGroupSize = 3
	group := make([][]byte, maxGroupSize)
	curGroupSize := 0
	for scanner.Scan() {
		group[curGroupSize] = getLine(scanner.Text())
		curGroupSize++

		if curGroupSize == maxGroupSize {
			curGroupSize = 0
			commonLetter, err := findCommonLetter(group)
			if err != nil {
				log.Panic(err)
			}

			if 'a' <= commonLetter && commonLetter <= 'z' {
				sum += int(commonLetter-'a') + 1
			} else {
				sum += int(commonLetter-'A') + 27
			}
		}
	}

	fmt.Println(sum)
}
