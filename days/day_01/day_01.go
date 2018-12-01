package main

import (
	"AdventOfCode/utils/disk"
	"AdventOfCode/utils/slice"
	"log"
)

func main() {
	const dataFilePath = "days/day_01/input.txt"
	frequencies := disk.GetIntSliceFromFile(dataFilePath)

	log.Printf("Puzzle №1: %d", Calibrate(frequencies))
	log.Printf("Puzzle №2: %d", CalibrateToFirstSeenTwice(frequencies))
}

func Calibrate(frequencies []int) int {
	return slice.Sum(frequencies)
}

func CalibrateToFirstSeenTwice(frequencies []int) (result int) {
	seen := make(map[int]bool)

	for {
		for _, frequency := range frequencies {
			if seen[result] {
				return result
			}

			seen[result] = true

			result += frequency
		}
	}
}
