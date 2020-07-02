package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	in.Scan()
	for i := 0; i < 4; i++ {
		fmt.Printf("%[1]v and %[1]T\n", in.Text())
	}
}
