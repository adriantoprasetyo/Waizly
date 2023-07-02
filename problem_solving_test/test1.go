package main

import (
	"fmt"
	"os"

	"waizly.com/adriantoprasetyo/problem_solving_test/libs"
)

func main() {
	var (
		minVal int
		maxVal int
	)

	arr := os.Args[1:]

	if len(os.Args) == 6 {
		for i := 0; i < 5; i++ {
			number := 0
			for x := 0; x < 5; x++ {
				if arr[i] != arr[x] {
					number += libs.Conv(arr[x])
				}
			}

			if minVal > number || minVal == 0 {
				minVal = number
			}

			if number > maxVal {
				maxVal = number
			}
		}
	} else {
		fmt.Printf("missing argument, input must 5 number on argument")
		return
	}

	fmt.Printf("%v %v", minVal, maxVal)
}
