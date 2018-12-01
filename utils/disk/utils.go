package disk

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFileHandler(path string) *os.File {
	fileHandler, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can't open file: %e", err)
	}

	return fileHandler
}

func GetIntSliceFromFile(path string) (result []int) {
	file := getFileHandler(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Can't convert string to number: %e", err)
		}

		result = append(result, num)
	}

	return result
}
