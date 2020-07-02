package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/inancgumus/screen"
)

const ball rune = 9918
const empty rune = 32

const (
	left  = -1
	right = 1
	down  = 1
	up    = -1
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("enter [width], [height]")
		return
	}
	wd, errW := strconv.Atoi(args[0])
	ht, errH := strconv.Atoi(args[1])
	if errW != nil || errH != nil || wd < 0 || ht < 0 {
		fmt.Println("enter width and height as positive integers")
		return
	}

	grid := make([][]bool, wd)

	for i := range grid {
		grid[i] = make([]bool, ht)
	}
	grid[1][1] = true
	velocity := [2]int{right, down}
	position := [2]int{0, 0}

	screen.Clear()
	buffer := make([]rune, 0, wd*(ht+1))
	for {
		screen.MoveTopLeft()
		buffer = buffer[:0]
	toExit:
		for x := range grid {
			for y := range grid[0] {
				if grid[x][y] == true {
					if x == wd-1 {
						velocity[0] = -1
					} else if x == 0 {
						velocity[0] = 1
					}
					if y == ht-1 {
						velocity[1] = -1
					} else if y == 0 {
						velocity[1] = 1
					}
					position[0], position[1] = x+velocity[0], y+velocity[1]
					grid[x][y] = false
					grid[position[0]][position[1]] = true
					if x == 0 && y == 0 || x == wd-1 && y == ht-1 || x == 0 && y == ht-1 || x == wd-1 && y == 0 {
						screen.Clear()
						fmt.Println("YOU WON !!!")
						time.Sleep(2 * time.Second)
						return
					}
					break toExit
				}
			}
		}
		for x := 0; x < wd; x++ {
			for y := 0; y < ht; y++ {
				if x == position[0] && y == position[1] {
					buffer = append(buffer, ball)
				} else {
					buffer = append(buffer, empty)
				}
			}
			buffer = append(buffer, '\n')
		}
		fmt.Print(string(buffer))
		time.Sleep(time.Millisecond * 50)
	}

}
