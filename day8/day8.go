package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Run() {
	num := 8
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", num))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println("Cannot converst char")
			}
			row = append(row, num)
		}
		grid = append(grid, row)
	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}

	results := make([][]int, len(grid))
	for i := range results {
		results[i] = make([]int, len(grid[i]))
	}

	// top
	//LEFT
	//for each row
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			results[i][j] = scenicScore(grid, i, j)
		}
	}

	max := 0
	for _, row := range results {
		for _, col := range row {
			if col > max {
				max = col
			}
		}
	}

	fmt.Println(fmt.Sprintf("\nDay %d: %d\n", num, max))

}

// next steps would be memoization
// really this is a dp problem, you can probably do the same solve from top left and bottom right again, each time calcuating the score for the tree? idk
func scenicScore(grid [][]int, row, col int) int {
	//scenic score is found by multiplying each of the directions number of trees scene...
	//look up
	up := grid[row][col]
	upSeen := 0
	for i := col - 1; i >= 0; i-- {
		upSeen++
		if grid[row][i] >= up {
			break
		}
	}

	//look right
	right := grid[row][col]
	rightSeen := 0
	for i := row + 1; i < len(grid); i++ {
		rightSeen++
		if grid[i][col] >= right {
			break
		}
	}

	//look down
	down := grid[row][col]
	downSeen := 0
	for i := col + 1; i < len(grid); i++ {
		downSeen++
		if grid[row][i] >= down {
			break
		}
	}

	//look left
	left := grid[row][col]
	leftSeen := 0
	for i := row - 1; i >= 0; i-- {
		leftSeen++
		if grid[i][col] >= left {
			break
		}
	}

	return upSeen * rightSeen * downSeen * leftSeen

}
