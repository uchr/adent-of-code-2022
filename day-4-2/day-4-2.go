package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type segment struct {
	l, r int
}

func parseSegment(text string) (segment, error) {
	n := strings.Split(text, "-")
	var err error
	var s segment
	s.l, err = strconv.Atoi(n[0])
	if err != nil {
		return s, err
	}
	s.r, err = strconv.Atoi(n[1])
	if err != nil {
		return s, err
	}

	return s, nil
}

func isSubset(s0, s1 segment) bool {
	return s0.l <= s1.l && s1.r <= s0.r
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}

	points := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		s0, err := parseSegment(parts[0])
		if err != nil {
			log.Panic(err)
		}
		s1, err := parseSegment(parts[1])
		if err != nil {
			log.Panic(err)
		}

		if isSubset(s0, s1) || isSubset(s1, s0) {
			points++
		}
	}

	fmt.Println(points)
}
