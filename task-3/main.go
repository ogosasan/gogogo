package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	fmt.Scanf("%d\r", &N)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	var a []string
	a = strings.Split(text, " ")
	var last = a[N-1]
	for i := N - 1; i > 0; i-- {
		a[i] = a[i-1]
	}
	a[0] = last
	fmt.Print(a)

}
