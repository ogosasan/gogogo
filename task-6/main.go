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
	var str []string
	var str_unic []string
	var count = 0
	str = strings.Split(text, " ")
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str_unic); j++ {
			if str[i] == str_unic[j] {
				count++
			}
		}
		if count == 0 {
			str_unic = append(str_unic, str[i])
		}
		count = 0
	}
	fmt.Print(len(str_unic))
}
