package main

import (
	"fmt"
	"testing"
)

func displaySL(sl *Skiplist) {
	for i := range sl.heads {
		cur := sl.heads[i]
		fmt.Printf("[%d]: ", i)
		for cur != nil {
			fmt.Printf("%d(%d)->", cur.value, cur.count)
			cur = cur.next
		}
		fmt.Printf("nil\n")
	}
}

// ["Skiplist","add","add","add","search","add","search","erase","erase","search"]
// [[],[1],[2],[3],[0],[4],[1],[0],[1],[1]]
// [null,null,null,null,false,null,true,false,true,false]
/*
func TestAnswer(t *testing.T) {
	skiplist := Constructor()

	skiplist.Add(1)
	displaySL(&skiplist)
	skiplist.Add(2)
	displaySL(&skiplist)
	skiplist.Add(3)
	displaySL(&skiplist)
	res := skiplist.Search(0) // return false.
	if res != false {
		t.Errorf("expect %v, but %v", false, res)
	}
	skiplist.Add(4)
	displaySL(&skiplist)
	res = skiplist.Search(1) // return true.
	if res != true {
		t.Errorf("expect %v, but %v", true, res)
	}
	res = skiplist.Erase(0) // return false, 0 is not in skiplist.
	displaySL(&skiplist)
	if res != false {
		t.Errorf("expect %v, but %v", false, res)
	}
	res = skiplist.Erase(1) // return true.
	displaySL(&skiplist)
	if res != true {
		t.Errorf("expect %v, but %v", true, res)
		fmt.Printf("expect %v, but %v\n", true, res)
	}
	res = skiplist.Search(1) // return false, 1 has already been erased.
	if res != false {
		t.Errorf("expect %v, but %v", false, res)
		fmt.Printf("expect %v, but %v\n", true, res)
	}
}
*/
/*
func TestAnswer2(t *testing.T) {
	skiplist := Constructor()

	skiplist.Add(10)
	displaySL(&skiplist)
	skiplist.Add(2)
	displaySL(&skiplist)
	skiplist.Add(3)
	displaySL(&skiplist)
	res := skiplist.Search(0) // return false.
	if res != false {
		t.Errorf("expect %v, but %v", false, res)
	}
	skiplist.Add(4)
	displaySL(&skiplist)
	res = skiplist.Search(1) // return true.
	if res != false {
		t.Errorf("expect %v, but %v", true, res)
	}
	res = skiplist.Erase(0) // return false, 0 is not in skiplist.
	displaySL(&skiplist)
	if res != false {
		t.Errorf("expect %v, but %v", false, res)
	}
	res = skiplist.Erase(10) // return true.
	displaySL(&skiplist)
	if res != true {
		t.Errorf("expect %v, but %v", true, res)
	}
	res = skiplist.Search(10) // return false, 1 has already been erased.
	if res != false {
		t.Errorf("expect %v, but %v", false, res)
	}
}
*/
func TestAnswer3(t *testing.T) {
	input1 := []string{"add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "add", "search", "add", "erase", "search", "erase", "search", "erase", "erase", "erase", "add", "erase", "add", "search", "erase", "search", "add", "erase", "erase", "erase", "add", "erase", "erase", "add", "erase", "add", "erase", "search", "erase", "search", "search", "erase", "search", "search", "add", "erase", "search", "search", "erase", "search", "add", "add", "search", "erase", "search", "search", "search", "add", "search", "add", "add", "add", "add", "search", "erase", "add", "search", "add", "search", "erase", "add", "add", "add", "erase", "search", "search", "search", "add", "add", "erase", "add", "erase", "add", "search", "add", "search", "search", "search", "search", "erase", "add", "erase", "search", "search", "search", "search", "erase", "search", "erase", "add", "add", "add", "search", "erase", "search", "search", "add", "add", "erase", "add", "erase", "search", "erase", "search", "erase", "add", "search", "search", "search", "search", "search", "search", "search", "search", "search", "search", "search", "search", "search", "search", "search"}
	input2 := []int{22, 25, 0, 11, 8, 1, 22, 3, 26, 15, 14, 3, 28, 17, 26, 21, 20, 11, 2, 17, 14, 9, 24, 13, 10, 29, 4, 25, 28, 7, 8, 29, 2, 9, 0, 23, 2, 3, 6, 9, 26, 1, 18, 13, 0, 15, 18, 7, 2, 9, 22, 1, 12, 13, 12, 15, 4, 19, 14, 21, 6, 3, 8, 1, 24, 1, 26, 27, 0, 21, 16, 27, 22, 23, 28, 5, 2, 9, 16, 1, 16, 23, 8, 5, 6, 19, 6, 27, 0, 21, 26, 5, 14, 25, 24, 21, 4, 21, 28, 25, 28, 23, 8, 27, 24, 1, 8, 17, 4, 21, 4, 19, 8, 23, 4, 11, 22, 25, 6, 9, 28, 11, 8, 25, 6, 5, 18, 29, 20, 15, 2, 3, 26, 15, 6}

	//expect := []bool{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, false, nil, true, false, true, true, false, false, false, nil, false, nil, true, true, false, nil, false, false, false, nil, false, false, nil, false, nil, true, true, false, false, true, true, false, false, nil, true, true, true, false, false, nil, nil, true, false, true, false, false, nil, true, nil, nil, nil, nil, false, true, nil, false, nil, true, false, nil, nil, nil, false, false, true, false, nil, nil, true, nil, false, nil, true, nil, true, true, false, true, true, nil, true, true, false, true, false, true, true, true, nil, nil, nil, true, false, true, false, nil, nil, true, nil, true, false, true, false, true, nil, false, false, true, false, false, false, false, false, true, true, true, true, true, true, false}
	expect := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, true, true, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, true, true, false, false, false, true, true, true, false, false, false, false, true, false, true, false, false, false, true, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, false, true, false, false, false, true, false, false, false, true, false, true, true, false, true, true, false, true, true, false, true, false, true, true, true, false, false, false, true, false, true, false, false, false, true, false, true, false, true, false, true, false, false, false, true, false, false, false, false, false, true, true, true, true, true, true, false}

	skiplist := Constructor()

	for i := range input1 {
		switch input1[i] {
		case "add":
			skiplist.Add(input2[i])
			displaySL(&skiplist)
		case "search":
			res := skiplist.Search(input2[i])
			if res != expect[i] {
				t.Errorf("expect %v, but %v", expect[i], res)
				fmt.Printf("!!expect %v, but %v\n", expect[i], res)
			}
		case "erase":
			res := skiplist.Erase(input2[i])
			if res != expect[i] {
				t.Errorf("expect %v, but %v", expect[i], res)
				fmt.Printf("!!expect %v, but %v\n", expect[i], res)
			}
			displaySL(&skiplist)
		}

	}

}

/*
func TestAnswer2(t *testing.T) {
	input := ""
	words := []string{}
	for i := 0; i < 5000; i++ {
		input = input + "a"
		words = append(words, "a")
	}
	//expect := 399980000
	_ = findSubstring(input, words)
	//if output != expect {
	//	t.Errorf("case: expect %v, but %v", expect, output)
	//}
}
*/
