package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("enter [text with http]")
		return
	}
	buff := make([]byte, 0, len(args[0]))
	for i, v := range args[0] {
		if v[i-7:i] == "http://" {
			buff = append(buff, "*"...)
		} else if (len(buff) != 0 && buff[i-1] == "*"[0]) && v != " " {
			buff = append(buff, "*"...)
		} else {
			buff = append(buff, v...)
		}
	}
	fmt.Println(buff)
}
