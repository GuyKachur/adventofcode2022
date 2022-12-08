package day6

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Run() {
	num := 6
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", num))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewReader(file)
	// unique := make(map[rune]int, 4)
	buff := make([]rune, 14)
	count := 0
	for {
		r, _, err := scanner.ReadRune()
		if err != nil {
			if err == io.EOF {
				fmt.Println("reached end of file without finding unique set")
				break
			}
			fmt.Println(err)
			break
		}
		count++
		buff = append(buff, r) // temp make buff 5
		buff = buff[1:]        // make buff the last 4
		if count > 14 {
			if unique(buff) {
				fmt.Println(fmt.Sprintf("\nDay %d: %d\n", num, count))
				return
			}
		}

	}

}

//so really i need a funciton that works on a slice... that can tell if theyre all unique

// buff will be constrained to be 4 long always so im cool with making it do 8 comparisons...
func unique(buff []rune) bool {
	for i, r := range buff {
		for j := i + 1; j < len(buff); j++ {
			if r == buff[j] {
				return false
			}
		}
	}
	return true
}
