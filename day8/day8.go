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

	//now that the grid is created lets find all the trees that are accesible
	//left
	// top, left, right, bottom := 0, 0, 0, 0
	canSee := make(map[string]bool, 0)
	results := make([][]int, len(grid))
	for i := range results {
		results[i] = make([]int, len(grid[i]))
	}

	// top
	//LEFT
	//for each row
	for i := 0; i < len(grid); i++ {
		leftMax := -1
		// go through each col an
		for j := 0; j < len(grid[i]); j++ {
			//if each cell can see the top mark as seen
			// update the max heigh weve seen
			if grid[i][j] > leftMax {
				leftMax = grid[i][j]
				canSee[fmt.Sprintf("%d:%d", i, j)] = true
				results[i][j] = 1
			}
		}
	}

	//TOP
	//for every col
	for j := 0; j < len(grid[0]); j++ {
		topMax := -1
		// go through every row
		for i := 0; i < len(grid); i++ {
			if grid[i][j] > topMax {
				topMax = grid[i][j]
				canSee[fmt.Sprintf("%d:%d", i, j)] = true
				results[i][j] = 1
			}
		}
	}

	//RIGHT
	//for each row
	for i := len(grid) - 1; i >= 0; i-- {
		rightMax := -1
		// go through each col an
		for j := len(grid[0]) - 1; j >= 0; j-- {
			//if each cell can see the top mark as seen
			// update the max heigh weve seen
			if grid[i][j] > rightMax {
				rightMax = grid[i][j]
				canSee[fmt.Sprintf("%d:%d", i, j)] = true
				results[i][j] = 1
			}
		}
	}

	//BOTTOM
	//for every col
	for j := len(grid[0]) - 1; j >= 0; j-- {
		bottomMax := -1
		// go through every row
		for i := len(grid) - 1; i >= 0; i-- {
			if grid[i][j] > bottomMax {
				bottomMax = grid[i][j]
				canSee[fmt.Sprintf("%d:%d", i, j)] = true
				results[i][j] = 1
			}
		}
	}

	//if i is row
	// and j is col
	// then as i go down a col i should keep track of top max
	// and as i go down a row keep track of leftmax

	total := 0
	for k, v := range canSee {
		if v {
			total++
		} else {
			fmt.Println(k)
		}
	}

	fmt.Println(fmt.Sprintf("\nDay %d: %d\n", num, total))

}
