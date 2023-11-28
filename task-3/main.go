package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d\r", &N)
	b := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&b[i])
	}
	var last = b[N-1]
	for i := N - 1; i > 0; i-- {
		b[i] = b[i-1]
	}
	b[0] = last
	fmt.Print(b)

}
