package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://adventofcode.com/2022/day/2

// The winner of the whole tournament is the player with the highest score. Your total score is the sum of your scores for each round. The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

// part two...
// x y z isnt what you throw its round outcome...
// x is a lose
// y is a draw
// z is win

//score is the same...
//round 1

// func round(opponent, mine string) int {
// 	score := 0
// 	switch mine {
// 	case "X":
// 		score += 1
// 	case "Y":
// 		score += 2
// 	case "Z":
// 		score += 3
// 	}
// 	if score == 1 {
// 		//i played rock
// 		switch opponent {
// 		case "A":
// 			score += 3 //they played rock so draw
// 		case "B":
// 			score += 0 // they played paper so lose
// 		case "C":
// 			score += 6 //they played scissors so win
// 		}
// 	} else if score == 2 {
// 		// i played paper
// 		switch opponent {
// 		case "A":
// 			score += 6
// 		case "B":
// 			score += 3
// 		case "C":
// 			score += 0
// 		}
// 	} else {
// 		// i played scissors
// 		switch opponent {
// 		case "A":
// 			score += 0
// 		case "B":
// 			score += 6
// 		case "C":
// 			score += 3
// 		}
// 	}

// 	return score
// }

func round(opponent, mine string) int {
	score := 0
	switch mine {
	case "X":
		score += 0
	case "Y":
		score += 3
	case "Z":
		score += 6
	}
	if score == 0 {
		//i need to lose
		switch opponent {
		case "A":
			score += 3 //they played rock and i need to lose so scissors.... which have a score of 3
		case "B":
			score += 1 // they played paper i need rock
		case "C":
			score += 2 //they played scissors so i need paper
		}
	} else if score == 3 {
		// i need to draw so... same as them
		switch opponent {
		case "A":
			score += 1
		case "B":
			score += 2
		case "C":
			score += 3
		}
	} else {
		// i need to win...
		switch opponent {
		case "A":
			score += 2
		case "B":
			score += 3
		case "C":
			score += 1
		}
	}

	return score
}

func Run() {
	num := 2
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", num))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), " ")
		total += round(lines[0], lines[1])
	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("\nDay %d: %d\n", num, total))
}
