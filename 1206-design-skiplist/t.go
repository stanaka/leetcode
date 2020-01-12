package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	value int
	count int
	next  *Node
	down  *Node
}

type Skiplist struct {
	levels int
	heads  []*Node
}

func Constructor() Skiplist {
	res := Skiplist{}
	res.levels = 4
	res.heads = make([]*Node, res.levels)
	for i := res.levels - 1; i >= 0; i-- {
		if i == res.levels-1 {
			res.heads[i] = &Node{value: -1, count: 1, next: nil, down: nil}
		} else {
			res.heads[i] = &Node{value: -1, count: 1, next: nil, down: res.heads[i+1]}
		}
	}
	return res
}

func (this *Skiplist) Search(target int) bool {
	//fmt.Printf("Search: %d\n", target)

	cur := this.heads[0]
	if cur == nil {
		return false
	}
	level := 0
	for {
		if target < cur.value {
			return false
		} else if cur.value == target {
			return true
		} else if cur.next != nil && cur.next.value <= target {
			cur = cur.next
		} else if cur.down != nil {
			cur = cur.down
			level++
		} else {
			break
		}
		////fmt.Printf("level: %d, value: %d\n", level, cur.value)
	}
	return false
}

func (this *Skiplist) Add(num int) {
	//fmt.Printf("Add: %d\n", num)
	if this.heads[0] == nil { // initialize
		for i := 0; i < this.levels; i++ {
			this.heads[i] = &Node{value: num, count: 1, next: nil, down: nil}
		}
		for i := 0; i < this.levels-1; i++ {
			this.heads[i].down = this.heads[i+1]
		}
		return
	}

	level := 0
	cur := make([]*Node, this.levels)
	prev := make([]*Node, this.levels)
	cur[level] = this.heads[level]
	for {
		if cur[level].value == num {
			break
		}
		if cur[level].next == nil {
			if cur[level].down != nil {
				cur[level+1] = cur[level].down
				level++
			} else { // cur[level].next == nil && cur[level].down == nil
				break
			}
		} else if cur[level].next.value <= num { // && cur[level].next != nil
			prev[level] = cur[level]
			cur[level] = cur[level].next
		} else { // cur[level].next != nil && cur[level].value < num < cur[level].next.value
			if cur[level].down != nil {
				cur[level+1] = cur[level].down
				level++
			} else { // cur[level].next != nil && cur[level].value < num < cur[level].next.value && cur[level].down == nil
				break
			}
		}
	}

	if cur[level].value == num {
		cur[level].count++
		for cur[level].down != nil {
			cur[level] = cur[level].down
			cur[level].count++
		}
		return
	}

	if cur[level].next == nil {
		if this.levels-1 == level {
			cur[level].next = &Node{value: num, count: 1, next: nil, down: nil}
			for rand.Int()%2 == 0 && level > 0 {
				level--
				cur[level].next = &Node{value: num, count: 1, next: nil, down: cur[level+1].next}
			}
			return
		} else {
			// shouldn't reach here
			fmt.Println("shouldn't reach here 1")
		}
	} else {
		// cur[level].value < num < cur[level].next.value
		if this.levels-1 == level {
			cur[level].next = &Node{value: num, count: 1, next: cur[level].next, down: nil}
			for rand.Int()%2 == 0 && level > 0 {
				level--
				cur[level].next = &Node{value: num, count: 1, next: cur[level].next, down: cur[level+1].next}
			}
			return
		} else {
			// shouldn't reach here
			fmt.Printf("shouldn't reach here 2: %d, %d\n", level, cur[level].value)
		}
	}
}

func (this *Skiplist) Erase(num int) bool {
	//fmt.Printf("Erase: %d\n", num)

	cur := make([]*Node, this.levels)
	prev := make([]*Node, this.levels)
	level := 0
	cur[level] = this.heads[0]
	if cur[level] == nil {
		return false
	}
	for {
		if num < cur[level].value {
			return false
		} else if cur[level].value == num {
			if cur[level].count > 1 {
				cur[level].count--
				for cur[level].down != nil {
					cur[level+1] = cur[level].down
					level++
					cur[level].count--
				}
				return true
			}
			if prev[level] != nil {
				prev[level].next = cur[level].next
				for prev[level].down != nil {
					////fmt.Printf("alevel: %d, value: %d\n", level, cur[level].value)
					cur[level+1] = cur[level].down
					prev[level+1] = prev[level].down
					level++
					for prev[level].next != cur[level] {
						prev[level] = prev[level].next
					}
					prev[level].next = cur[level].next
				}
				return true
			}
			for level < this.levels {
				this.heads[level] = this.heads[level].next
				level++
			}
			return true
		} else if cur[level].next != nil && cur[level].next.value <= num {
			prev[level] = cur[level]
			cur[level] = cur[level].next
		} else if cur[level].down != nil {
			cur[level+1] = cur[level].down
			level++
		} else { // if cur[level].next == nil && cur[level].down == nil {
			return false
		}
		////fmt.Printf("level: %d, value: %d\n", level, cur[level].value)
	}
	return false
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
