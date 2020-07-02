package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	words := make(map[string]bool)
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		text := in.Text()
		text = strings.ToLower(text)
		if words[text] {
			fmt.Println("TWICE!!!")
			return
		} else {
			words[text] = true
		}
	}
}
