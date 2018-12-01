package main

import (
	"AdventOfCode/utils/disk"
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

	for _, k := range splitData {
		num, err := strconv.Atoi(k)
		if err != nil {
			log.Fatalf("Can't convert: %e", err)
		}

		result += num
	}

	log.Printf("Result is %d", result)
}
