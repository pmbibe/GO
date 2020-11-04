package main

import (
	"fmt"
	"math"
)

var max int = math.MaxInt64

func backtracking(arrS []string, s string, result string, count int) int {
	if len(s) == 0 {
		if count < max {
			max = count
		}
		return count
	}
	for i := 0; i < len(s)+1; i++ {
		part := s[0:i]

		for _, v := range arrS {
			if part == v {
				backtracking(arrS, string(s[i:len(s)]), result+""+part, count+1)
			}
		}

	}
	return max

}

func main() {
	arrS := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa"}
	s := "aaaaaa"
	fmt.Println(backtracking(arrS, s, "", 0))

}
