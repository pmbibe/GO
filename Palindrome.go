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

func symmetricPrime(n int) int {
	count := 0
	for m := 1; ; m++ {
		if KiemtraNT(m) == 1 && KiemtraNT(reverse(m)) == 1 && m == reverse(m) {
			count++
			if count == n {
				return m
			}

		}

	}
	return 0
}
func main() {

	fmt.Print(symmetricPrime(300))

}
