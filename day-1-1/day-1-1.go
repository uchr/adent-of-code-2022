package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}

	scanner := bufio.NewScanner(file)
	curCallory := 0
	maxCallory := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if maxCallory < curCallory {
				maxCallory = curCallory
			}
			curCallory = 0
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Panic(err)
		}
		curCallory += num
	}

	fmt.Println(maxCallory)
}
