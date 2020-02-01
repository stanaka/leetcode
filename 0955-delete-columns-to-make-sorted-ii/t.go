package main

//var indices []bool
//var flags []bool

func minDeletionSize(a []string) int {
	if len(a) == 0 {
		return 0
	}
	//fmt.Printf("%v\n", a)
	flags := make([]bool, len(a))
	indices := make([]bool, len(a[0]))

	for i := range a[0] {
		tmp := make([]bool, len(a))
		for j := 1; j < len(a); j++ {
			if flags[j] == true {
				continue
			}
			if a[j-1][i] < a[j][i] {
				tmp[j] = true
			} else if a[j-1][i] == a[j][i] {
			} else {
				indices[i] = true
				break
			}
		}
		if indices[i] == false {
			for j := range tmp {
				if tmp[j] == true {
					flags[j] = true
				}
			}
		}
	}

	ans := 0
	//fmt.Printf("%v\n", indices)
	for i := range indices {
		if indices[i] == true {
			ans++
		}
	}
	return ans
}

/*
func detect(a []string, i1, i2, j1, j2 int) {
	fmt.Printf("d: %d-%d, %d-%d\n", j1, j2, i1, i2)
	for i := i1; i < i2; i++ {
		if indices[i] == true {
			continue
		}
		flag := 0
		prev := uint8('a') - 1
		for j := j1; j < j2; j++ {
			if prev < a[j][i] {
				prev = a[j][i]
			} else if prev == a[j][i] {
				flag = 1
				js := j
				for j < j2 && prev == a[j][i] {
					j++
				}
				if i+1 < i2 {
					detect(a, i+1, i2, js-1, j)
				}
				if j < j2 {
					prev = a[j][i]
				}
			} else {
				flag = 2
				break
			}
		}
		fmt.Printf("[%d], flag:%d\n", i, flag)
		switch flag {
		case 2:
			indices[i] = true
		case 0:
			return
		case 1:
			return
		}
	}
}

func minDeletionSizeIncomplete(a []string) int {
	if len(a) == 0 {
		return 0
	}
	indices = make([]bool, len(a[0]))
	fmt.Printf("%v\n", a)
	detect(a, 0, len(a[0]), 0, len(a))
	ans := 0
	fmt.Printf("%v\n", indices)

	for i := range indices {
		if indices[i] == true {
			ans++
		}
	}
	return ans
}

*/
