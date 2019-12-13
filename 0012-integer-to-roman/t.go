package main

//import "fmt"

//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

type t struct {
	k int
	v string
}

func intToRoman(num int) string {
	ans := ""
	m := []t{
		{1000, "M"},
		{900 ,"CM"},
		{500 ,"D"},
		{400 ,"CD"},
		{100 ,"C"},
		{90,"XC"},
		{50,"L"},
		{40,"XL"},
		{10,"X"},
		{9,"IX"},
		{5,"V"},
		{4,"IV"},
		{1,"I"},
	}

	for _,v := range m {
		for num >= v.k {
			ans += v.v
			num -= v.k
		}
	}

	return ans
}
