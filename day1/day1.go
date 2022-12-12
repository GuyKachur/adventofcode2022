package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// https://adventofcode.com/2022/day/1
func Run() {
	num := 0
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", num))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total, podium := 0, max(0, 0, 0, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			podium = max(total, podium[0], podium[1], podium[2])
			total = 0
		} else {
			num, _ := strconv.Atoi(line)
			total += num
		}
	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}

	max := 0
	for i := 0; i < len(podium); i++ {
		// fmt.Println(podium[i])
		max += podium[i]
	}

	fmt.Println(fmt.Sprintf("\nDay %d: %d\n", num, max))

}

func max(one, two, three, four int) []int {
	arr := []int{one, two, three, four}
	sort.Ints(arr)
	return arr[1:]
}
