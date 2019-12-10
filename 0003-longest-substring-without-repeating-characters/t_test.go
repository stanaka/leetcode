package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	input := "abcabcbb"
	output := lengthOfLongestSubstring(input)
	if output != 3 {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase2(t *testing.T) {
	input := "bbbbb"
	output := lengthOfLongestSubstring(input)
	if output != 1 {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase3(t *testing.T) {
	input := "pwwkew"
	output := lengthOfLongestSubstring(input)
	if output != 3 {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase4(t *testing.T) {
	input := " "
	output := lengthOfLongestSubstring(input)
	if output != 1 {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase5(t *testing.T) {
	input := "abc"
	output := lengthOfLongestSubstring(input)
	if output != 3 {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase6(t *testing.T) {
	input := "dvdf"
	output := lengthOfLongestSubstring(input)
	if output != 3 {
		t.Errorf("Error!! %v", output)
	}
}
