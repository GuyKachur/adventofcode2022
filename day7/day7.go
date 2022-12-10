package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dirMap map[string]*myDir

func init() {
	// dirMap = make(map[string]*myDir, 0)
	// memo = make(map[string]int, 0)
}

type Node struct {
	isFile bool
	size   int
	name   string
	//if its a directory it has children
	children map[string]*Node
	parent   *Node
}

func (n Node) createChild(child *Node) error {
	if n.isFile {
		return fmt.Errorf("File unable to have children")
	}

	// updated := false
	// if node, ok := n.children[child.name]; ok {
	// 	//child is in map... update sizes
	// 	if node.size != child.size {
	// 		n.size = n.size - node.size + child.size
	// 	}

	// } else {
	//node isnt in children map
	n.children[child.name] = child
	n.size += child.size

	// }

	//if size was updated... i need to carry that upwards in the chain

	//updated sizes up the chain
	// if updated {
	var cursor *Node
	cursor = child.parent
	for cursor != nil {
		cursor.size = cursor.size + child.size
		cursor = cursor.parent
	}
	// }

	return nil
}

//functions to traverse nodes like tree?
// iterate over?

func Run() {
	num := 7
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", num))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// stack := []string{}
	// myNewDir := &myDir{
	// 	name:  "/",
	// 	files: make(map[*myFile]bool, 0),
	// 	dirs:  make(map[*myDir]bool, 0),
	// }
	// dirMap["/"] = myNewDir
	// fileMap := make(map[string]*myFile, 0)

	//setup root node

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
	// stack := []*Node{root}
	pointer := root

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")

		//
		//first check for cd,
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
		// then check for dir
		// then check for numbers

		// 	// if its a cd then i want to push the directory into the stack
		// 	if strings.HasPrefix(line, "$ cd") {
		// 		if words[2] == ".." {
		// 			stack = stack[:len(stack)-1]
		// 		} else {
		// 			//special case cd /... restarts the stack from empty
		// 			if words[2] == "/" {
		// 				stack = []string{"/"}
		// 			} else {
		// 				// we can append the directory to the stack
		// 				stack = append(stack, words[2])
		// 			}
		// 		}
		// 	} else if strings.HasPrefix(line, "$ ls") {
		// 		// this is the ls case... we just skip it?

		// 	} else {
		// 		// if the line starts with dir or numbers
		// 		if strings.HasPrefix(line, "dir") {
		// 			//words 1 is the directory name
		// 			dirName := words[1]
		// 			//get it from the map if its already there
		// 			var directory *myDir
		// 			if dir, ok := dirMap[dirName]; ok {
		// 				directory = dir
		// 			} else {
		// 				myNewDir := &myDir{
		// 					name:  dirName,
		// 					files: make(map[*myFile]bool, 0),
		// 					dirs:  make(map[*myDir]bool, 0),
		// 				}
		// 				directory = myNewDir
		// 			}
		// 			dirMap[dirName] = directory
		// 			//working directory is at top of stack
		// 			pwd := stack[len(stack)-1]
		// 			// so i want to check if the pwd is in the dirmap
		// 			if mypwd, ok := dirMap[pwd]; ok {
		// 				mypwd.dirs[directory] = true
		// 				dirMap[pwd] = mypwd
		// 			}

		// 		} else {
		// 			// its a file
		// 			sizeString, name := words[0], words[1]
		// 			size, err := strconv.Atoi(sizeString)
		// 			if err != nil {
		// 				fmt.Println(err)
		// 				os.Exit(1)
		// 			}
		// 			myFile := &myFile{
		// 				size: size,
		// 				name: name,
		// 			}
		// 			//working directory is at top of stack
		// 			pwd := stack[len(stack)-1]
		// 			// so i want to check if the pwd is in the dirmap
		// 			if mypwd, ok := dirMap[pwd]; ok {
		// 				mypwd.files[myFile] = true
		// 				mypwd.size += myFile.size
		// 				dirMap[pwd] = mypwd
		// 			} else {
		// 				fmt.Println("Error finding pwd in stack")
		// 			}
		// 			// if its not in the directory map... weve done something wrong

		// 			// then traverse up the stack and update each of those dir sizes
		// 			// starting at the second to last
		// 			for i := len(stack) - 2; i >= 0; i-- {
		// 				pwd := stack[i]
		// 				if mypwd, ok := dirMap[pwd]; ok {
		// 					mypwd.size += myFile.size
		// 					dirMap[pwd] = mypwd
		// 				}
		// 			}
		// 		}

		// 	}

	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// parsed everything into the dirs

	// assuming the map is correct....
	// what would a program to traverse the map look like

	// totalSize := countFileSizes("/")
	// total := 0
	// for _, v := range dirMap {
	// 	if v.size <= 100000 {
	// 		total += v.size
	// 	}
	// }
	total := bfs(root)

	fmt.Println(fmt.Sprintf("\nDay %d: %d\n", num, total))

}

//so the directory itself can be under... meaning nested directories count...

func bfs(root *Node) int {
	//make a queue... as long as that queue isnt done youre not done
	q := []*Node{root}
	total := 0
	for len(q) != 0 {
		pop := q[0]
		q = q[1:]
		if pop.isFile {
			continue
		}
		if pop.size <= 100000 {
			//add to total, move on
			total += pop.size
			//we dont add children here because weve accounted for them

		}
		// else {
		// if its over we want to add its children
		for _, v := range pop.children {
			q = append(q, v)
		}
		// }
	}
	return total
}

var memo map[string]int

func countFileSizes(root string) int {
	//so each directory size == total of file sizes + total of directories

	size := 0
	if dir, ok := dirMap[root]; ok {
		for file := range dir.files {
			size += file.size
		}
		for dirs := range dir.dirs {
			if answer, ok := memo[dirs.name]; ok {
				size += answer
			} else {
				size += countFileSizes(dirs.name)
			}
		}
		dir.size = size
		dirMap[root] = dir
	}
	return size
}

// // updates parents for filesizes
// func updateFileSizes(dirs *map[string]*myDir) {
// 	// could maybe alternatively go back up the stack and update each of those
// }

type myFile struct {
	size int
	name string
}

type myDir struct {
	name  string
	size  int //represents teh size of all lower directories
	files map[*myFile]bool
	dirs  map[*myDir]bool
}

// func parseLine(line string) {

// }

// i feel like i can like... just create a map
// if every time i find a size and then file... i can log that...
// every time i find a list.. i can make a pointer to each of those... eventually the root pointer will remain?
// this feels like a stack thing... find a cd push to stack... find a ls pull the directory off the stack and use that...
// what would cd do... pup the stack?

// goal is to find all the directories with total size of 100,000
//then calculate sum of total sizes
// find every directory with a size under 100k and sum their sizes

// instead of doing the like.. parse and whatnot
// i think i wanna create teh filesyetme like a madman
// and create a map of dir name to a list of files/sizes in it and a lits of pointers to othe ritems

func RunV2() {

	// var line string

}
