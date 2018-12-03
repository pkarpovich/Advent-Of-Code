package main

import "testing"

func TestFindMaxClaimsSize(t *testing.T) {
	test1 := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}
	ans1 := 2

	tests := []struct {
		testCase []string
		answer   int
	}{
		{test1, ans1},
	}

	for index, test := range tests {
		if ans := FindMaxClaimsSize(test.testCase); test.answer != ans {
			t.Errorf("[%d] got %d, want %d", index, ans, test.answer)
		}
	}
}
