package main

import (
	"fmt"
	"log"
	"os"
)

func findSubset(line []byte, windowSize int) int {
	indexOffset := windowSize - 1

	repeated := 0
	chars := make([]byte, 255)
	for i, c := range line {
		chars[c]++
		if chars[c] > 1 {
			repeated++
		}
		if i < indexOffset {
			continue
		}

		if repeated == 0 {
			return i + 1
		}

		chars[line[i-indexOffset]]--
		if chars[line[i-indexOffset]] != 0 {
			repeated--
		}
	}

	return -1
}

func main() {
	line, err := os.ReadFile("input")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(findSubset(line, 4))
	fmt.Println(findSubset(line, 14))
}
