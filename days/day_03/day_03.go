package main

import (
	"AdventOfCode/utils/disk"
	"fmt"
	"log"
)

func main() {
	const inputPath = "days/day_03/input.txt"
	inputData := disk.GetStringSliceFromFile(inputPath)

	claims := parseClaims(inputData)

	log.Printf("Puzzle №1: %d", FindMaxClaimsSize(claims))
}

type Claim struct {
	id   int
	x, y int
	w, h int
}

func (c *Claim) apply(grid [][]int) {
	for x := 0; x < c.w; x++ {
		for y := 0; y < c.h; y++ {
			grid[c.x+x][c.y+y]++
		}
	}
}

func (c *Claim) overlap(grid [][]int) bool {
	for x := 0; x < c.w; x++ {
		for y := 0; y < c.h; y++ {
			if grid[c.x+x][c.y+y] > 1 {
				return true
			}
		}
	}

	return false
}

func generateGrid(width, height int) [][]int {
	grid := make([][]int, width)
	for i := range grid {
		grid[i] = make([]int, height)
	}

	return grid
}

func parseClaims(input []string) (claims []Claim) {
	for _, inputLine := range input {
		var claim Claim

		_, err := fmt.Sscanf(
			inputLine,
			"#%d @ %d,%d: %dx%d",
			&claim.id,
			&claim.x,
			&claim.y,
			&claim.w,
			&claim.h,
		)
		if err != nil {
			log.Printf("Can't parse string '%s': %e", inputLine, err)
		}

		claims = append(claims, claim)
	}

	return claims
}

func FindMaxClaimsSize(claims []Claim) (result int) {
	grid := generateGrid(1000, 1000)

	for _, c := range claims {
		c.apply(grid)
	}

	for x := range grid {
		for _, cell := range grid[x] {
			if cell > 1 {
				result++
			}
		}
	}

	for _, claim := range claims {
		if !claim.overlap(grid) {
			fmt.Println(claim.id)
		}
	}

	return result
}
