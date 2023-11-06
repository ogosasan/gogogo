package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	var a []string
	a = strings.Split(text, "")
	var N = "one"
	for i := 0; i < len(a); i++ {
		if a[i] == "1" {
			a[i] = N
		}
	}
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i])
	}
}
