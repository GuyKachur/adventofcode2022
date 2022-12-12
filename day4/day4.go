package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	num := 4
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", num))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		var times = make([][]int, 0)
		lines := strings.Split(scanner.Text(), ",")
		for _, line := range lines {
			timestrings := strings.Split(line, "-")
			start, _ := strconv.Atoi(timestrings[0])
			end, _ := strconv.Atoi(timestrings[1])
			times = append(times, []int{start, end})
		}
		if fullyOverlap(times[0], times[1]) {
			total += 1
		}

	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("\nDay %d: %d\n", num, total))

}
func fullyOverlap(a, b []int) bool {
	start1, end1 := a[0], a[1]
	start2, end2 := b[0], b[1]

	// Check if one interval starts after the other ends
	if start1 > end2 || start2 > end1 {
		return false
	}

	//fully overlapping
	// // Check if one interval fully contains the other
	// if start1 >= start2 && end1 <= end2 {
	// 	return true
	// } else if start2 >= start1 && end2 <= end1 {
	// 	return true
	// }

	return true
}
