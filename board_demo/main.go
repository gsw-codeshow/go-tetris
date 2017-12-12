package main

import (
	"go-tetris/board"

	"github.com/nsf/termbox-go"
)

func main() {
	Board := board.Board{30, 10, 30, 10}
	termbox.Init()
	for i := 0; i < Board.Height; i++ {
		for j := 0; j < Board.Width; j++ {
			termbox.SetCell(i, j, rune(' '), termbox.ColorWhite, termbox.ColorWhite)
		}
	}
	termbox.Flush()
	for {

	}
}
