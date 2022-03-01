package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var res int
	for i := 0; i <= 9*N; i++ {
		temp := T(N, i)
		res += temp * temp
	}
	fmt.Print(res)
}

func T(n, s int) int {
	if n == 1 {
		return 1
	}
	var res int
	for i := s - 9; i <= s; i++ {
		if i < 0 || i > 9*(n-1) {
			continue
		}
		res += T(n-1, i)
	}
	return res
}
//aaa
