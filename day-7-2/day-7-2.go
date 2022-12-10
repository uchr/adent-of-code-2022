package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node interface {
	Parent() *FolderNode
	Size() int
}

type FolderNode struct {
	parent   *FolderNode
	children map[string]Node
}

func NewFolder(parent *FolderNode) *FolderNode {
	return &FolderNode{parent: parent, children: make(map[string]Node)}
}

func (n FolderNode) Parent() *FolderNode {
	return n.parent
}

func (n FolderNode) Size() int {
	folderSize := 0

	for _, node := range n.children {
		folderSize += node.Size()
	}

	return folderSize
}

type FileNode struct {
	parent *FolderNode
	size   int
}

func NewFile(parent *FolderNode, size int) *FileNode {
	return &FileNode{parent: parent, size: size}
}

func (n FileNode) Parent() *FolderNode {
	return n.parent
}

func (n FileNode) Size() int {
	return n.size
}

type System struct {
	root   *FolderNode
	curDir *FolderNode
}

func (s *System) parseList(reader *bufio.Reader) {
	for {
		lineStart, err := reader.Peek(1)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Panic(err)
		}
		if lineStart[0] == '$' {
			return
		}

		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Panic(err)
		}
		entry := strings.Split(strings.TrimSpace(line), " ")
		if entry[0] == "dir" {
			s.curDir.children[entry[1]] = NewFolder(s.curDir)
		} else {
			size, err := strconv.Atoi(entry[0])
			if err != nil {
				log.Panic()
			}
			s.curDir.children[entry[1]] = NewFile(s.curDir, size)
		}
	}
}

func (s *System) parseChangeDir(line string) {
	path := strings.Split(strings.TrimSpace(line), " ")[2]
	if path == ".." {
		s.curDir = s.curDir.parent
		return
	}

	folders := strings.Split(path, "/")

	startIndex := 0
	if folders[0] == "" {
		s.curDir = s.root
		startIndex = 1
	}

	for i := startIndex; i < len(folders); i++ {
		for folderName, folder := range s.curDir.children {
			if folderName == folders[i] {
				var ok bool
				s.curDir, ok = folder.(*FolderNode)
				if !ok {
					log.Panic(errors.New("failed cast folder"))
				}
				break
			}
		}
	}
}

func (s *System) parseCommands(reader *bufio.Reader) {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Panic(err)
		}

		if strings.HasPrefix(line, "$ ls") {
			s.parseList(reader)
		} else if strings.HasPrefix(line, "$ cd") {
			s.parseChangeDir(line)
		}
	}
}

func printSize(nodeName string, node Node, prefix string) {
	fmt.Printf("%s %s %d\n", prefix, nodeName, node.Size())

	if folderNode, ok := node.(*FolderNode); ok {
		for childName, child := range folderNode.children {
			printSize(childName, child, prefix+"-")
		}
	}
}

func getFolderSizes(node Node, sizes *[]int) {
	if folderNode, ok := node.(*FolderNode); ok {
		*sizes = append(*sizes, node.Size())

		for _, child := range folderNode.children {
			getFolderSizes(child, sizes)
		}
	}
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}

	reader := bufio.NewReader(file)
	s := System{root: NewFolder(nil)}
	s.parseCommands(reader)
	printSize("/", s.root, "")

	folderSizes := make([]int, 0)
	getFolderSizes(s.root, &folderSizes)

	sort.Ints(folderSizes)

	curFreeSpace := 70000000 - folderSizes[len(folderSizes)-1]
	minFolderSize := 30000000 - curFreeSpace

	fmt.Println("curFreeSpace", curFreeSpace)
	fmt.Println("minFolderSize", minFolderSize)

	for i := len(folderSizes) - 2; i >= 0; i-- {
		if folderSizes[i] < minFolderSize {
			fmt.Println("Answer:", i, folderSizes[i+1])
			break
		}
	}
}
