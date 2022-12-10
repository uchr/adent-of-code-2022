package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func updateTop(topCallories []int, callory int) {
	replacedValue := -1
	replaceIndex := -1
	for i := 0; i < len(topCallories); i++ {
		if topCallories[i] >= callory {
			continue
		}

		replacedValue = topCallories[i]
		topCallories[i] = callory
		replaceIndex = i
		break
	}

	if replaceIndex < 0 {
		return
	}

	for i := replaceIndex + 1; i < len(topCallories); i++ {
		temp := topCallories[i]
		topCallories[i] = replacedValue
		replacedValue = temp
	}
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}

	scanner := bufio.NewScanner(file)
	curCallory := 0

	topCallories := make([]int, 3)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			updateTop(topCallories, curCallory)
			curCallory = 0
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Panic(err)
		}
		curCallory += num
	}

	sum := 0
	for _, v := range topCallories {
		sum += v
	}

	fmt.Println(sum)
}
