package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("inputSmall")
	if err != nil {
		log.Panic(err)
	}

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []byte(scanner.Text())

		sort.Slice(line[:len(line)/2], func(i, j int) bool { return line[i] < line[j] })
		sort.Slice(line[len(line)/2:], func(i, j int) bool { return line[len(line)/2+i] < line[len(line)/2+j] })

		var commonLetter byte
		for li, ri := 0, len(line)/2; li < len(line)/2 && ri < len(line); {
			if line[li] == line[ri] {
				commonLetter = line[li]
				break
			}

			if line[li] < line[ri] {
				li++
			} else {
				ri++
			}
		}

		if 'a' <= commonLetter && commonLetter <= 'z' {
			sum += int(commonLetter-'a') + 1
		} else {
			sum += int(commonLetter-'A') + 27
		}
	}

	fmt.Println(sum)
}
