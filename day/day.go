package day

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	num := 0
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", num))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		total += len(line)
	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("\nDay %d: %d\n", num, total))

}
