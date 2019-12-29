package main

//import "fmt"

// TrieNode ...
type TrieNode struct {
	Num   int
	Count int
	Next  map[rune]*TrieNode
}

// Track ...
type Track struct {
	T     *TrieNode
	Pos   int
	Flag  map[int]int
	Count int
}

func findChild(r *TrieNode, c rune) *TrieNode {
	if v, ok := r.Next[c]; ok {
		return v
	}
	return nil
}

func createTrie(words *[]string) (*TrieNode, map[int]int) {
	t := &TrieNode{0, 0, map[rune]*TrieNode{}}
	res := map[int]int{}
	for index, w := range *words {
		i := t
		for _, c := range w {
			n := findChild(i, c)
			if n != nil {
				i = n
			} else {
				i.Next[c] = &TrieNode{0, 0, map[rune]*TrieNode{}}
				i = i.Next[c]
			}
		}
		if i.Count == 0 {
			i.Num = index
			i.Count = 1
		} else {
			i.Count++
		}
		res[i.Num] = i.Count
	}
	return t, res
}

func findSubstring(s string, words []string) []int {
	ans := []int{}

	root, resultMap := createTrie(&words)
	//fmt.Printf("resultMap: %v\n", resultMap)

	/*fmt.Printf("%v\n", *root)
	for _, n := range root.Next {
		fmt.Printf("%v\n", *n)
		for _, n := range n.Next {
			fmt.Printf("%v\n", *n)
			for _, n := range n.Next {
				fmt.Printf("%v\n", *n)
			}
		}
	}*/
	tracks := []*Track{}

	lastPos := len(s)
	if len(words) > 0 {
		lastPos -= (len(words) - 1) * len(words[0])
	}
	for p, r := range s {
		for _, t := range tracks {
			if t.T == nil {
				continue
			}
			i := findChild(t.T, r)
			if i != nil {
				t.T = i
				if i.Count != 0 {
					t.Flag[i.Num]++
					t.Count++
					if t.Flag[i.Num] > resultMap[i.Num] || t.Count == len(words) {
						t.T = nil
					} else {
						t.T = root
					}
				}
			} else {
				t.T = nil
			}
		}
		if p > lastPos {
			continue
		}
		i := findChild(root, r)
		if i != nil {
			if i.Count == 0 {
				tracks = append(tracks, &Track{i, p, make(map[int]int), 0})
			} else {
				t := &Track{nil, p, make(map[int]int), 0}
				if i.Count != 0 {
					t.Flag[i.Num]++
					t.Count++
					if t.Flag[i.Num] > resultMap[i.Num] || t.Count == len(words) {
						t.T = nil
					} else {
						t.T = root
					}
				}
				tracks = append(tracks, t)
			}
		}
		//for _, n := range tracks {
		//	fmt.Printf("[%d]%d: %v\n", p, r, n)
		//}
	}

	for _, t := range tracks {
		flag := true
		for k, v := range resultMap {
			if f, ok := t.Flag[k]; !ok || f != v {
				flag = false
			}
		}
		if flag == true {
			ans = append(ans, t.Pos)
		}
	}
	return ans
}

func main() {
	return
}
