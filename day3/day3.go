package day3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// first half represent the first compartment
// second half reps the second compartment...
// compartments should have all the same type of items...
// so im looking for the one character thats in both the first and second one...
// they also have priority... 1-26 for a-z and 27-52 for A-Z
// sum all the priorities
func Run() {
	// lets say we have that boilerplate code for reading in lines...

	num := 3
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", num))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()
		cnt := true
		for _, char := range line {
			if cnt {
				if strings.Contains(line2, string(char)) {
					if strings.Contains(line3, string(char)) {
						if char > 90 {
							total -= 96
						} else {
							total -= 38
						}
						total += int(char)
						cnt = false
					}
				}
			}
		}
	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("\nDay %d: %d\n", num, total))

}
