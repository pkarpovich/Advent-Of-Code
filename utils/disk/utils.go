package disk

import (
	"io/ioutil"
	"log"
)

func GetFileStringData(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Can't open file: %e", err)
	}

	return string(data)
}
