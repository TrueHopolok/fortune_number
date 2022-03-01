package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var res uint64
	for i := 0; i <= 9*n; i++ {
		temp := searcher(n, i)
		res += temp * temp
	}
	fmt.Print(res)
}

func searcher(n, s int) uint64 {
	if n == 1 {
		return 1
	}
	var res uint64
	for i := s - 9; i <= s; i++ {
		if i < 0 || i > 9*(n-1) {
			continue
		}
		res += searcher(n-1, i)
	}
	return res
}
