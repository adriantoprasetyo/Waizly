package main

import (
	"fmt"
	"strings"

	"waizly.com/adriantoprasetyo/problem_solving_test/libs"
)

func main() {
	var (
		arrLen  int
		minNum  int
		zeroNum int
		plusNum int
	)

	arrLen = libs.Conv(libs.ScanInput().Text())
	arr := strings.Split(libs.ScanInput().Text(), " ")

	if arrLen == len(arr) {
		for _, element := range arr {
			number := libs.Conv(element)

			if number < 0 {
				minNum++
			} else if number == 0 {
				zeroNum++
			} else if number > 0 {
				plusNum++
			}
		}
		fmt.Printf(" %.6f \n %6f \n %6f", float32(plusNum)/float32(arrLen), float32(minNum)/float32(arrLen), float32(zeroNum)/float32(arrLen))
	} else {
		fmt.Printf("missing argument, total number does not match array length")
		return
	}
}
