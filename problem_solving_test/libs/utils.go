package libs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ScanInput() *bufio.Scanner {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input
}

func Conv(input string) (number int) {
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("error parsing argument, %v", err)
		return
	}
	return number
}
