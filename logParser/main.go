package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Log Parser from Stratch
//
//  You've watched the lecture. Now, try to create the same
//  log parser program on your own. Do not look at the lecture,
//  and the existing source code.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < log.txt
//
//   DOMAIN                             VISITS
//   ---------------------------------------------
//   blog.golang.org                        30
//   golang.org                             10
//   learngoprogramming.com                 20
//
//   TOTAL                                  60
//
//
//  go run main.go < log_err_missing.txt
//
//   wrong input: [golang.org] (line #3)
//
//
//  go run main.go < log_err_negative.txt
//
//   wrong input: "-100" (line #3)
//
//
//  go run main.go < log_err_str.txt
//
//   wrong input: "FOUR" (line #3)
//
// ---------------------------------------------------------

func main() {
	in := bufio.NewScanner(os.Stdin)
	log := make(map[string]int)
	lop := make([]string, 0, 20)
	var total int
	for in.Scan() {
		str := in.Text()
		site, n := strings.Fields(str)[0], strings.Fields(str)[1]
		visit, _ := strconv.Atoi(n)
		log[site] += visit
		total += visit
		lop = append(lop, site)
	}
	fmt.Printf("%-20s %20s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 40))
	for _, v := range lop {
		fmt.Printf("%-38s %d\n", v, log[v])
	}
	fmt.Printf("%-20s %20d\n", "TOTAL", total)
}
