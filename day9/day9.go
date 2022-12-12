package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	num := 9
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", num))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var link *Leader
	tail := &Leader{
		self:      &Pointer{x: 0, y: 0, visited: &map[string]int{}, moves: make([]string, 0)},
		followers: nil,
	}
	(*tail.self.visited)[fmt.Sprintf("%d:%d", tail.self.x, tail.self.y)] = 1
	link = tail
	for i := 0; i < 9; i++ {
		link = &Leader{
			self:      &Pointer{x: 0, y: 0, visited: &map[string]int{}, moves: make([]string, 0)},
			followers: []*Leader{link},
		}
		(*link.self.visited)[fmt.Sprintf("%d:%d", link.self.x, link.self.y)] = 1
	}
	head := link

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		moves, err := strconv.Atoi(words[1])
		if err != nil {
			fmt.Printf("Unable to tell num moves")
			os.Exit(1)
		}
		for i := 0; i < moves; i++ {
			switch words[0] {
			case "D":
				head.Down()
			case "U":
				head.Up()
			case "R":
				head.Right()
			case "L":
				head.Left()
			default:
				fmt.Println("How did you get here")
			}
		}

	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("\nDay %d: %d\n", num, len(*tail.self.visited)))

}
