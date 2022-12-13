package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	nodes := make([]ListNode, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		nodes = append(nodes, *ParseNode(line))
	}

	delimeter2 := ListNode{nodes: []Node{ListNode{nodes: []Node{ValueNode{value: 2}}}}}
	delimeter6 := ListNode{nodes: []Node{ListNode{nodes: []Node{ValueNode{value: 6}}}}}

	nodes = append(nodes, delimeter2)
	nodes = append(nodes, delimeter6)

	sort.Slice(nodes, func(i, j int) bool { return nodes[i].Compare(nodes[j]) == -1 })

	index2, index6 := 0, 0
	for i := range nodes {
		if nodes[i].Compare(delimeter2) == 0 {
			index2 = i + 1
			fmt.Println(i)
		} else if nodes[i].Compare(delimeter6) == 0 {
			index6 = i + 1
			fmt.Println(i)
		}
		fmt.Println(nodes[i])
	}

	fmt.Println(index2 * index6)
}
