package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	total, words := 0, make(map[string]int)
	rx := regexp.MustCompile(`[^a-z('s)]`)
	for in.Scan() {
		total++
		word := strings.ToLower(in.Text())
		word = rx.ReplaceAllString(word, "")
		words[word]++
	}
	fmt.Println(total, len(words))
}
