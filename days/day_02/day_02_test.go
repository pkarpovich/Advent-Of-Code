package main

import "testing"

func TestChecksum(t *testing.T) {
	test1 := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	ans1 := 12

	tests := []struct {
		testCase []string
		answer   int
	}{
		{test1, ans1},
	}

	for index, test := range tests {
		if ans := Checksum(test.testCase); test.answer != ans {
			t.Errorf("[%d] got %d, want %d", index, ans, test.answer)
		}
	}
}

func TestFindFabricBoxes(t *testing.T) {
	test1 := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}
	ans1 := "fgij"

	tests := []struct {
		testCase []string
		answer   string
	}{
		{test1, ans1},
	}

	for index, test := range tests {
		if ans := FindFabricBoxes(test.testCase); test.answer != ans {
			t.Errorf("[%d] got '%s', want '%s'", index, ans, test.answer)
		}
	}
}
