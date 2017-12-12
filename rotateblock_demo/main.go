package main

import (
	"fmt"
	"go-tetris/block"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	bf := block.InitBlockInterface()
	termboxErr := termbox.Init()
	if nil != termboxErr {
		fmt.Println(termboxErr)
	}
	currentShapes := bf.Entry[1]
	currentX := 10
	currentY := 10
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		currentShapes = bf.Rotate(currentShapes)
		for _, shapesItem := range currentShapes.Body {
			termbox.SetCell(shapesItem.X+currentX, shapesItem.Y+currentY, rune('#'), termbox.ColorGreen, termbox.ColorDefault)
		}
		termbox.Flush()
		time.Sleep(time.Second)
	}
}
