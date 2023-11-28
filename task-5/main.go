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
	for i := 0; i < len(a); i++ {
		a[i] = strings.ReplaceAll(a[i], "1", "one")
	}
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i])
	}
}
