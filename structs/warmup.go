package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}
type export struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
}

var encoded []byte

func main() {
	games := []game{
		{item: item{id: 1, name: "god of war", price: 50}, genre: "action adventure"},
		{item{2, "x-com 2", 30}, "strategy"},
		{item{3, "minecraft", 20}, "sandbox"},
	}
	byID := make(map[int]game)
	for _, v := range games {
		byID[v.id] = v
	}
	fmt.Println()

	in := bufio.NewScanner(os.Stdin)
	const commands string = `Available commands are:
		list   -> list all games
		id nbr -> game with id nbr
		save   -> encode into json & quit
		quit   -> quits to prompt
	`
	fmt.Println(commands)
	for in.Scan() {
		if in.Text() == "" {
			return
		}
		text := in.Text()

		if text == "list" {
			fmt.Printf("Aybat's game store has %d games.\n\n", len(games))

			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}
		} else if text == "quit" {
			return
		} else if strings.Contains(text, "id") {
			textList := strings.Fields(text)
			if len(textList) != 2 {
				fmt.Println("wrong id")
				goto here
			}
			id, err := strconv.Atoi(textList[1])
			if err != nil {
				fmt.Println("wrong id")
				goto here
			}
			if game, ok := byID[id]; ok {
				fmt.Printf("%-3s %-12s %-17s %s\n", "id", "title", "genre", "price")
				fmt.Printf("#%-2d %-12s %-17s $%d\n", game.id, game.name, game.genre, game.price)
			} else {
				fmt.Println("sorry, we don't have this item")
			}
		} else if text == "save" {
			forExport := make([]export, len(games), 10)
			for i, v := range games {
				forExport[i].ID = v.id
				forExport[i].Genre = v.genre
				forExport[i].Name = v.name
				forExport[i].Price = v.price
			}
			encoded, err := json.MarshalIndent(forExport, "", "\t")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("\n", string(encoded))
		} else if text == "decode" {
			var decoded []export
			fmt.Println(decoded)
			fmt.Println([]byte(data))
			err := json.Unmarshal([]byte(data), &decoded)
			if err != nil {
				return
			}
			fmt.Println(decoded)
		}
	here:
		fmt.Println()
	}

}
