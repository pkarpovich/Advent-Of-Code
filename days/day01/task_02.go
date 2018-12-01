package main

import (
	"AdventOfCode/utils/disk"
	"AdventOfCode/utils/slice"
	"log"
	"strconv"
	"strings"
)

func main() {
	const separator = "\r\n"
	const dataFilePath = "days/day01/data.txt"
	fileData := disk.GetFileStringData(dataFilePath)

	splitData := strings.Split(fileData, separator)

	var result int
	var arrayResults []int

	for i := 0; true; i++ {
		if i == len(splitData) {
			i = 0
		}

		numStr := splitData[i]

		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalf("Can't convert string to int: %s", err)
		}

		result += num

		if slice.Contains(arrayResults, result) {
			log.Printf("Contains result: %d", result)
			return
		}

		arrayResults = append(arrayResults, result)
	}
}
