package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	isFile   bool
	size     int
	name     string
	children map[string]*Node
	parent   *Node
}

func (n Node) createChild(child *Node) error {
	if n.isFile {
		return fmt.Errorf("File unable to have children")
	}

	n.children[child.name] = child
	n.size += child.size

	var cursor *Node
	cursor = child.parent
	for cursor != nil {
		cursor.size = cursor.size + child.size
		cursor = cursor.parent
	}

	return nil
}

func Run() {
	num := 7
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", num))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	root := &Node{
		size:     0,
		name:     "root",
		isFile:   false,
		children: make(map[string]*Node, 0),
		parent:   nil,
	}
	slash := &Node{
		size:     0,
		name:     "/",
		isFile:   false,
		children: make(map[string]*Node, 0),
		parent:   root,
	}
	root.children[slash.name] = slash
	pointer := root

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")

		if strings.HasPrefix(line, "$ cd") {
			if words[2] == ".." {
				// stack = stack[1:]
				pointer = pointer.parent
			} else {
				// move pointer down into the matching child
				if child, ok := pointer.children[words[2]]; ok {
					pointer = child
				} else {
					fmt.Println("Unable to find pointer in child node")
				}

			}
		} else if strings.HasPrefix(line, "dir ") {
			//create the dir we see as a child to this directory
			newFile := &Node{
				size:     0,
				name:     words[1],
				isFile:   false,
				children: make(map[string]*Node, 0),
				parent:   pointer,
			}
			pointer.createChild(newFile)

		} else {
			//ls... or file size
			if !strings.HasPrefix(line, "$ ls") {
				//file processing
				sizeString, name := words[0], words[1]
				size, err := strconv.Atoi(sizeString)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				newFile := &Node{
					size:     size,
					name:     name,
					isFile:   true,
					children: nil,
					parent:   pointer,
				}
				pointer.createChild(newFile)
			}
		}

	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}

	total := bfs(root)
	fmt.Println(fmt.Sprintf("\nDay %d: %d\n", num, total))

}

func bfs(root *Node) int {
	q := []*Node{root}
	total := root.size
	unused := 70000000 - total
	small := root
	for len(q) != 0 {
		pop := q[0]
		q = q[1:]
		if pop.isFile {
			continue
		}
		//part 1
		// if pop.size <= 100000 {
		// 	//add to total, move on
		// 	total += pop.size
		// }

		//part 2
		//smallest node that canfree up enough space to run update
		if unused+pop.size > 30000000 {
			// deleting this node could free up enough space
			if pop.size < small.size {
				small = pop
			}
		}

		for _, v := range pop.children {
			q = append(q, v)
		}

	}
	return small.size
}
