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

	file, err := os.Open("day3/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// so aline is going to be a rucksack
		//group
		line := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()
		// halves DmptngtF   |||  wvvMmwmm
		// i want to check each character?
		// mid := len(line) / 2
		// comp1, comp2 := line[:mid], line[mid:]
		cnt := true
		for _, char := range line {
			if cnt {
				if strings.Contains(line2, string(char)) {
					// so if char is below 91 then we can subtract 38..
					// if char is above... we subtract 96
					if strings.Contains(line3, string(char)) {
						// the third contains it too weve found our char
						if char > 90 {
							total -= 96
						} else {
							total -= 38
						}
						total += int(char)
						// fmt.Println(string(char))
						cnt = false
					}
				}
			}
		}
	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(total)

	// then once ive found the char thats there.. add that to priorities

}
