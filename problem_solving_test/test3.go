package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"waizly.com/adriantoprasetyo/problem_solving_test/libs"
)

func main() {
	var (
		timestamp  string
		hourFormat string
	)

	timeString := os.Args[1]
	timestamp = timeString[0 : len(timeString)-2]
	hourFormat = timeString[len(timeString)-2:]
	timeResult := strings.Split(timestamp, ":")

	formatCheck(timeResult)

	if hourFormat == "AM" {

		if timeResult[0] == "12" {
			timeResult[0] = "00"
		}
	} else if hourFormat == "PM" {
		number := libs.Conv(timeResult[0])
		number += 12
		timeResult[0] = strconv.Itoa(number)
	} else {
		fmt.Printf("unknown time format")
		return
	}

	fmt.Printf("%v:%v:%v \n", timeResult[0], timeResult[1], timeResult[2])
}

func formatCheck(data []string) {
	for i := 0; i < len(data); i++ {
		libs.Conv(data[i])
	}
}
