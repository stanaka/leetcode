package main

import "testing"

func TestCase1(t *testing.T) {
	input1 := []int{1, 3}
	input2 := []int{2}
	expected := 2.0
	output := findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
	output = findMedianSortedArrays(input2, input1)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase2(t *testing.T) {
	input1 := []int{1, 2}
	input2 := []int{3, 4}
	expected := 2.5
	output := findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
	output = findMedianSortedArrays(input2, input1)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase3(t *testing.T) {
	input1 := []int{1, 3}
	input2 := []int{3, 4}
	expected := 3.0
	output := findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
	output = findMedianSortedArrays(input2, input1)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase4(t *testing.T) {
	input1 := []int{1, 2, 3}
	input2 := []int{4}
	expected := 2.5
	output := findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
	output = findMedianSortedArrays(input2, input1)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase5(t *testing.T) {
	input1 := []int{1, 4, 7}
	input2 := []int{2, 4, 6}
	expected := 4.0
	output := findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
	output = findMedianSortedArrays(input2, input1)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase6(t *testing.T) {
	input1 := []int{}
	input2 := []int{1}
	expected := 1.0
	output := findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
	output = findMedianSortedArrays(input2, input1)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase7(t *testing.T) {
	input1 := []int{}
	input2 := []int{1, 2}
	expected := 1.5
	output := findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
	output = findMedianSortedArrays(input2, input1)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase8(t *testing.T) {
	input1 := []int{1, 2}
	input2 := []int{-1, 3}
	expected := 1.5
	output := findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
	output = findMedianSortedArrays(input2, input1)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase9(t *testing.T) {
	input1 := []int{1}
	input2 := []int{2, 3}
	expected := 2.0
	output := findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
	output = findMedianSortedArrays(input2, input1)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase10(t *testing.T) {
	input1 := []int{2}
	input2 := []int{1, 3, 4}
	expected := 2.5
	output := findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
	output = findMedianSortedArrays(input2, input1)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase11(t *testing.T) {
	input1 := []int{3}
	input2 := []int{1, 2, 4}
	expected := 2.5
	output := findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
	output = findMedianSortedArrays(input2, input1)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase12(t *testing.T) {
	input1 := []int{3}
	input2 := []int{1, 2, 4, 5}
	expected := 3.0
	output := findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
	output = findMedianSortedArrays(input2, input1)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
}

func TestCase13(t *testing.T) {
	input1 := []int{5}
	input2 := []int{1, 2, 3, 4, 6}
	expected := 3.5
	output := findMedianSortedArrays(input1, input2)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
	output = findMedianSortedArrays(input2, input1)
	if output != expected {
		t.Errorf("Error!! %v", output)
	}
}
