package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d\t", &n)
	arr := make([][]string, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		var a []string
		a = strings.Split(text, " ")
		for j := 0; j < n; j++ {
			arr[i][j] = a[j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == arr[j][i] {
				continue
			} else {
				fmt.Println("no")
				return
			}
		}
	}
	fmt.Println("yes")
}
