package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Read() []int {
	input := bufio.NewScanner(os.Stdin)
	var numbers []int

	for input.Scan() {
		line := strings.TrimSpace(input.Text())

		if line == "q" || line == "Q" {
			break
		}

		num, err := strconv.Atoi(line)

		if num <= -100000 || num >= 100000 {
			fmt.Println("The entered value is greater than 100000 or less than -100000")
			continue
		} else if err != nil {
			fmt.Println("Invalid input data:", err)
			continue
		} else {
			numbers = append(numbers, num)
		}
	}
	return numbers
}
