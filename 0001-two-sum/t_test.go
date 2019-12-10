package main

import (
	"testing"
)

type testCase struct {
	input  string
	expect string
}

func TestTwoSum1(t *testing.T) {
	output := twoSum([]int{2, 7, 11, 15}, 9)
	if len(output) != 2 || output[0] != 0 || output[1] != 1 {
		t.Errorf("Error!! %v", output)
	}

}

func TestTwoSum2(t *testing.T) {
	output := twoSum([]int{3, 2, 4}, 6)
	if len(output) != 2 || output[0] != 1 || output[1] != 2 {
		t.Errorf("Error!! %v", output)
	}

}
