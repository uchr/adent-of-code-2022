package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type Node interface {
	Compare(other Node) int
}

type ValueNode struct {
	value int
}

func (v ValueNode) Compare(node Node) int {
	switch o := node.(type) {
	case ValueNode:
		if v.value < o.value {
			return -1
		} else if v.value > o.value {
			return 1
		} else {
			return 0
		}
	case ListNode:
		l := ListNode{nodes: []Node{v}}
		return l.Compare(node)
	}

	return 0
}

func (v ValueNode) String() string {
	return fmt.Sprintf("%v", v.value)
}

type ListNode struct {
	nodes []Node
}

func (l ListNode) Compare(node Node) int {
	switch o := node.(type) {
	case ValueNode:
		ol := ListNode{nodes: []Node{node}}
		return l.Compare(ol)
	case ListNode:
		i, j := 0, 0
		for {
			if i == len(l.nodes) && j == len(o.nodes) {
				return 0
			} else if i == len(l.nodes) && j != len(o.nodes) {
				return -1
			} else if i != len(l.nodes) && j == len(o.nodes) {
				return 1
			}

			if cmpResult := l.nodes[i].Compare(o.nodes[j]); cmpResult != 0 {
				return cmpResult
			}

			i++
			j++
		}
	}

	return 0
}

func (l ListNode) String() string {
	sb := strings.Builder{}
	sb.WriteByte('[')
	for i := 0; i < len(l.nodes); i++ {
		sb.WriteString(fmt.Sprintf("%v", l.nodes[i]))
		if i != len(l.nodes)-1 {
			sb.WriteByte(',')
		}
	}
	sb.WriteByte(']')
	return sb.String()
}

func ParseNode(line string) *ListNode {
	stack := []ListNode{}
	var topStack *ListNode

	curNumber := 0
	isNumber := false

	for i := 0; i < len(line); i++ {
		if line[i] == '[' {
			stack = append(stack, ListNode{})
			topStack = &stack[len(stack)-1]
		}

		if (line[i] == ',' || line[i] == ']') && isNumber {
			topStack.nodes = append(topStack.nodes, ValueNode{value: curNumber})
			curNumber = 0
			isNumber = false
		}

		if line[i] == ']' {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				stack[len(stack)-1].nodes = append(stack[len(stack)-1].nodes, *topStack)
				topStack = &stack[len(stack)-1]
			}
		}

		if unicode.IsNumber(rune(line[i])) {
			d := int(line[i] - '0')
			curNumber = curNumber*10 + d
			isNumber = true
		}
	}

	return topStack
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}

	scanner := bufio.NewScanner(file)
	pairs := make([]ListNode, 2)
	pairIndex := 0
	index := 1
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		pairs[pairIndex] = *ParseNode(line)
		pairIndex++

		if pairIndex == 2 {
			pairIndex = 0
			fmt.Println(pairs[0])
			fmt.Println(pairs[1])
			fmt.Println(pairs[0].Compare(pairs[1]))
			if pairs[0].Compare(pairs[1]) < 0 {
				sum += index
			}
			index++
		}
	}

	fmt.Println(sum)

}
