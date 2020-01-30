package main

import "fmt"

func minDistance(w1 string, w2 string) int {
	//fmt.Printf("%v\n", w1)
	//fmt.Printf("%v\n", w2)
	b := make([]int, len(w1))

	for i := range w2 {
		max := 0
		for j := range w1 {
			t := max
			if t < b[j] {
				t = b[j]
			}
			if w2[i] == w1[j] {
				b[j] = max + 1
			}
			max = t
		}
		//fmt.Printf("%v\n", b)
	}
	max := 0
	for j := range w1 {
		if max < b[j] {
			max = b[j]
		}
	}
	return len(w1) + len(w2) - max*2
}

func minDistance_array(word1 string, word2 string) int {
	b := make([][]int, 26)
	a := make([][][]int, len(word2)+1)
	for i := range b {
		b[i] = []int{}
	}

	for i := range word1 {
		c := word1[i] - 'a'
		b[c] = append(b[c], i)
	}

	a[0] = [][]int{[]int{0, 0}}
	for i := range word2 {
		a[i+1] = [][]int{}
		c := word2[i] - 'a'
		if len(b[c]) > 0 {
			for k := range a[i] {
				a[i+1] = append(a[i+1], a[i][k])
				for j := range b[c] {
					if a[i][k][0] < b[c][j] {
						a[i+1] = append(a[i+1], []int{b[c][j], a[i][k][1] + 1})
						break
					}
				}
			}
		} else {
			a[i+1] = a[i]
		}
	}
	max := 0
	if len(a) > 0 {
		for i := range a[len(a)-1] {
			if max < a[len(a)-1][i][1] {
				max = a[len(a)-1][i][1]
			}
		}

		fmt.Printf("%v\n", a)
		return len(word1) - max + len(word2) - max
	}
	return len(word1)
}
