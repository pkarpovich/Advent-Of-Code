package main

import "testing"

func TestCalibrate(t *testing.T) {
	test1 := []int{1, 1, 1}
	ans1 := 3

	test2 := []int{1, 1, -2}
	ans2 := 0

	test3 := []int{-1, -2, -3}
	ans3 := -6

	tests := []struct {
		testCase []int
		answer   int
	}{
		{test1, ans1},
		{test2, ans2},
		{test3, ans3},
	}

	for index, test := range tests {
		if ans := Calibrate(test.testCase); test.answer != ans {
			t.Errorf("[%d] got %d, want %d", index, ans, test.answer)
		}
	}
}

func TestCalibrateToFirstSeenTwice(t *testing.T) {
	test1 := []int{1, -1}
	ans1 := 0

	test2 := []int{3, 3, 4, -2, -4}
	ans2 := 10

	test3 := []int{-6, 3, 8, 5, -6}
	ans3 := 5

	test4 := []int{7, 7, -2, -7, -4}
	ans4 := 14

	tests := []struct {
		testCase []int
		answer   int
	}{
		{test1, ans1},
		{test2, ans2},
		{test3, ans3},
		{test4, ans4},
	}

	for index, test := range tests {
		if ans := CalibrateToFirstSeenTwice(test.testCase); ans != test.answer {
			t.Errorf("[%d] got %d, want %d", index, ans, test.answer)
		}
	}
}
