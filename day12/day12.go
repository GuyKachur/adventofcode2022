package day12

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	num := 12
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", num))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]rune
	lc := 0
	start := []int{}
	end := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, make([]rune, len(line)))
		for i, char := range line {
			if char == 'S' {
				start = []int{lc, i}
				char = 'a'
			} else if char == 'E' {
				end = []int{lc, i}
				char = 'z'
			}
			grid[lc][i] = char
		}
		lc++
	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}

	shortest := 999999
	for i, row := range grid {
		for j, val := range row {
			if val == 'a' {
				res := bfs([]int{i, j}, end, &grid)
				if res < shortest && res > 0 {
					shortest = res
					start = []int{i, j}
				}
			}
		}
	}
	fmt.Println(fmt.Sprintf("\nDay %d: %d from [%d,%d]\n", num, shortest, start[0], start[1]))

}

func bfs(start, end []int, grid *[][]rune) int {
	visited := make([][]int, len((*grid)))
	for i := range *grid {
		visited[i] = make([]int, len((*grid)[i]))
	}

	seen := make([][]int, len((*grid)))
	for i := range *grid {
		seen[i] = make([]int, len((*grid)[i]))
	}

	q := make([][]int, 0)
	q = append(q, start)
	count := 0
	wave := len(q)
	for len(q) != 0 {
		pop := q[0]
		q = q[1:]
		x, y := pop[0], pop[1]
		if num := visited[x][y]; num == 0 || num > count {
			visited[x][y] = count
		}

		char := (*grid)[x][y]
		if x == end[0] && y == end[1] {
			return visited[x][y]
		}

		//up
		if visitable(char, x-1, y, &visited, grid) {
			if seen[x-1][y] != 1 {
				q = append(q, []int{x - 1, y})
				seen[x-1][y] = 1
			}
		}
		//right
		if visitable(char, x, y+1, &visited, grid) {
			if seen[x][y+1] != 1 {
				q = append(q, []int{x, y + 1})
				seen[x][y+1] = 1
			}
		}
		//down
		if visitable(char, x+1, y, &visited, grid) {
			if seen[x+1][y] != 1 {
				q = append(q, []int{x + 1, y})
				seen[x+1][y] = 1
			}
		}
		//left
		if visitable(char, x, y-1, &visited, grid) {
			if seen[x][y-1] != 1 {
				q = append(q, []int{x, y - 1})
				seen[x][y-1] = 1
			}
		}
		wave--
		if wave == 0 {
			count++
			wave = len(q)
		}
	}
	return visited[end[0]][end[1]]
}

func visitable(char rune, x, y int, visited *[][]int, grid *[][]rune) bool {
	if x < len((*visited)) && x >= 0 {
		//x is in board
		if y < len((*visited)[0]) && y >= 0 {
			//y is also in board
			if ok := (*visited)[x][y]; ok == 0 {
				decide := (*grid)[x][y] - char
				if decide <= 1 {
					return true
				}
			}
		}
	}
	return false
}

// part two should be done the other way around...
// should search from the end to the beginning... the visitable rules would need to be reformatted
// then whenever a char of 'a' is found, we can record the depth for it... and compare to the smallest
// then we only have to do the search once instead of... every time we find an a....
