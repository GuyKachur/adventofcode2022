package day5

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	num := 5
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", num))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pic := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			pic = append(pic, line)
		} else {
			break
		}
	}
	// fields is fine for the key line... but not the others
	cols := len(strings.Fields(pic[len(pic)-1]))
	stacks := make([][]string, cols)
	for i := range stacks {
		stacks[i] = make([]string, 0)
	}

	// so the stacks will be the last item,
	// and then for the other rows, starting from the bottom
	// add each letter to the stacks
	re := regexp.MustCompile(`\[[A-Z]\]| {3} |   `)
	for i := len(pic) - 2; i >= 0; i-- {
		idx := cols - 1
		line := pic[i]
		//going backwards
		matches := re.FindAllString(line, -1)

		for j := len(matches) - 1; j >= 0; j-- {
			if matches[j] != "    " {
				stacks[idx] = append(stacks[idx], matches[j])
			}
			idx--
		}
	}

	for scanner.Scan() {
		// line := scanner.Text()
		lines := strings.Split(scanner.Text(), " ")
		limitString, startString, destString := lines[1], lines[3], lines[5]
		limit, _ := strconv.Atoi(limitString)
		start, _ := strconv.Atoi(startString)
		dest, _ := strconv.Atoi(destString)
		//o index
		start--
		dest--

		// start becomes start without the last limit elements
		// dest becomes dest with the last three elements appended
		startStack := stacks[start]
		crane := startStack[len(startStack)-limit:]
		destStack := append(stacks[dest], crane...)

		stacks[dest] = destStack
		//then start can just have those last elements shaved off
		stacks[start] = startStack[:len(stacks[start])-limit]

	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}
	answer := ""
	for _, stack := range stacks {
		answer += stack[len(stack)-1]
	}
	fmt.Println(fmt.Sprintf("\nDay %d: %s\n", num, answer))

}

// package main

// import (
//   "fmt"
//   "strings"
// )

// // ParseInput takes a string representing the input and returns a slice of
// // strings, where each string is a column of the input with the bottommost
// // letter first.
// func ParseInput(input string) []string {
//   // Split the input into individual lines
//   lines := strings.Split(input, "\n")

//   // Find the maximum number of columns in any row
//   maxCols := 0
//   for _, line := range lines {
//     numCols := len(strings.Fields(line))
//     if numCols > maxCols {
//       maxCols = numCols
//     }
//   }

//   // Create a slice of strings to hold the columns
//   cols := make([]string, maxCols)

//   // Loop through the lines in reverse order
//   for i := len(lines) - 1; i >= 0; i-- {
//     // Split the line into individual words
//     words := strings.Fields(lines[i])

//     // Loop through the words
//     for j, word := range words {
//       // Append the word to the appropriate column
//       cols[j] += word
//     }
//   }

//   return cols
// }

// func main() {
//   input := `
// [T]     [D]         [L]
// [R]     [S] [G]     [P]         [H]
// [G]     [H] [W]     [R] [L]     [P]
// [W]     [G] [F] [H] [S] [M]     [L]
// [Q]     [V] [B] [J] [H] [N] [R] [N]
// [M] [R] [R] [P] [M] [T] [H] [Q] [C]
// [F] [F] [Z] [H] [S] [Z] [T] [D] [S]
// [P] [H] [P] [Q] [P] [M] [P] [F] [D]
//   1   2   3   4   5   6   7   8   9
// `

//   cols := ParseInput(input)
//   for _, col := range cols {
//     fmt.Println(col)
//   }
// }
