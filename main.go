package main

import "fmt"

func main() {
	var N uint64
	fmt.Scan(&N)
	var res uint64
	for i := uint64(0); i <= 9*N; i++ {
		temp := T(N, i)
		res += temp * temp
	}
	fmt.Print(res)
}

func T(n, s uint64) uint64 {
	if n == 1 {
		return 1
	}
	var res uint64
	for i := s - 9; i <= s; i++ {
		if i < 0 || i > 9*(n-1) {
			continue
		}
		res += T(n-1, i)
	}
	return res
}
