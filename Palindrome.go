package main

import (
	"fmt"
	"math"
)

func KiemtraNT(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	m := math.Sqrt(float64(n))
	for i := 2; i <= int(m); i++ {
		if int(n)%i == 0 {
			return 0
		}

	}

	return 1
}

func reverse(n int) int {
	m := 0
	for n > 0 {
		m = m*10 + n%10
		n = n / 10
	}
	return m
}

func main() {
	count := 0

	for n := 1; n < 300; n++ {
		if KiemtraNT(n) == 1 && KiemtraNT(reverse(n)) == 1 {
			count++
			if count == 1 {
				fmt.Print(n)
			}

		}

	}

}
