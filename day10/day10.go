package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type crt struct {
	row, col int
	board    [][]bool
}

func (c *crt) Mark(x int) {
	if x+1 == c.col || x-1 == c.col || x == c.col {
		c.board[c.row][c.col] = true
	} else {
		c.board[c.row][c.col] = false
	}
}

func (c *crt) Next(x int) {
	c.Mark(x)
	if c.col+1 == 40 {
		c.col = 0
		if c.row+1 == 6 {
			c.row = 0
		} else {
			c.row += 1
		}
	} else {
		c.col += 1
	}
}

func Run() {
	num := 10
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", num))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	x := 1
	cycle := 0
	cycles := []int{x}
	board := make([][]bool, 6)
	for i := 0; i < 6; i++ {
		board[i] = make([]bool, 40)
	}
	cursor := &crt{
		row:   0,
		col:   0,
		board: board,
	}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "noop") {
			cycle += 1
			cycles = append(cycles, x)
			cursor.Next(x)
		} else {
			//addx
			cycle++
			cycles = append(cycles, x)
			cursor.Next(x)

			words := strings.Split(line, " ")
			numStr := words[1]
			neg := false
			if strings.HasPrefix(numStr, "-") {
				neg = true
				numStr = numStr[1:]
			}
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Unable to parse number string: " + numStr)
			}
			cycle++
			cycles = append(cycles, x)
			cursor.Next(x)
			if neg {
				x -= num
			} else {
				x += num
			}

		}
	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}

	forty := 20
	sum := 0
	for ind, num := range cycles {
		if ind == forty {
			sum += (num * ind)
			forty += 40
		}
	}

	fmt.Println(fmt.Sprintf("\nDay %d:%d\n", num, sum))

	// for _, row := range cursor.board {
	// 	for _, boo := range row {
	// 		if boo {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }
}
