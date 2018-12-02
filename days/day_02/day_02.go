package main

import (
	"AdventOfCode/utils/disk"
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"log"
	"strings"
)

func main() {
	const dataFilePath = "days/day_02/input.txt"
	ids := disk.GetStringSliceFromFile(dataFilePath)

	log.Printf("Puzzle №1: %d", Checksum(ids))
	log.Printf("Puzzle №2: %s", FindFabricBoxes(ids))
}

func Checksum(ids []string) (result int) {
	const two = 2
	const three = 3

	idsInfo := map[int]int{}

	for _, id := range ids {
		idChars := make(map[int]bool)

		for _, char := range id {
			count := strings.Count(id, string(char))
			strings.Replace(id, string(char), "", -1)

			if idChars[two] && idChars[three] {
				break
			}

			if idChars[count] {
				continue
			}

			idChars[count] = true

			idsInfo[count]++
		}
	}

	return idsInfo[two] * idsInfo[three]
}

type FabricBox struct {
	first    string
	second   string
	maxRatio float64
}

func FindFabricBoxes(ids []string) (result string) {
	fb := FabricBox{}

	for _, first := range ids {
		for _, second := range ids {
			if first == second {
				continue
			}

			ratio := levenshtein.RatioForStrings(
				[]rune(first),
				[]rune(second),
				levenshtein.DefaultOptions,
			)

			if fb.maxRatio < ratio {
				fb = FabricBox{
					first,
					second,
					ratio,
				}
			}
		}
	}

	log.Printf("First: %s", fb.first)
	log.Printf("Second: %s", fb.second)
	log.Printf("Ratio: %f", fb.maxRatio)

	for index, char := range fb.first {
		if fb.second[index] == fb.first[index] {
			result += string(char)
		}
	}

	return result
}
