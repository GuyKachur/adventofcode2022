package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {

	// 71-97,71-72
	// 60-97,20-95
	// 20-59,58-59

	file, err := os.Open("day4/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		//fix input
		var times = make([][]int, 0)
		lines := strings.Split(scanner.Text(), ",")
		//71-97
		for _, line := range lines {
			timestrings := strings.Split(line, "-")
			start, _ := strconv.Atoi(timestrings[0])
			end, _ := strconv.Atoi(timestrings[1])
			times = append(times, []int{start, end})
		}
		// so we have an int[] of int[]s
		// and im trying to count how many pairs have an overlap
		//overlap
		if fullyOverlap(times[0], times[1]) {
			total += 1
		}

	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(total)

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
