package main

import (
	"fmt"
)

const text = `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.

Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.

Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

const (
	sp         = ' '
	tab        = '\t'
	newLine    = '\n'
	lineLength = 40 - 1
)

var runes []rune

func main() {
	args := text
	if len(args) == 0 {
		fmt.Println("enter text")
		return
	}
	var prevSpace int
	var counter int
	for i, v := range args {
		if v == ' ' {
			prevSpace = i
		}
		if v == '\n' {
			counter = 0
		}
		if counter > lineLength {
			if runes[len(runes)-1] == ' ' || runes[len(runes)-1] == '\t' {
				runes = append(runes[:len(runes)-1], '\n')
				counter = 0
			} else if runes[len(runes)-1] == '\n' {
				counter = 0
			} else {
				var lefts []rune
				lefts = append(lefts, runes[len(runes)-(i-prevSpace):]...)
				runes = append(runes[:len(runes)-(i-prevSpace)], '\n')
				runes = append(runes, lefts...)
				counter = 0
			}
		}
		counter++
		runes = append(runes, v)
	}
	fmt.Println(string(runes))
}
