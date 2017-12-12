package main

import (
	"go-tetris/block"

	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type Coord struct {
	X, Y int
}

func main() {
	termbox.Init()
	defer termbox.Close()

	bf := block.InitBlockInterface()
	bfEnum := len(bf.Entry)
	for {
		bfEntryIndex := rand.Int() % bfEnum
		currentShapes := bf.Entry[bfEntryIndex]
		currentCoord := Coord{10, 0}
		for _, currentShapesCoord := range currentShapes.Body {
			if currentCoord.X < bf.Abs(currentShapesCoord.X) {
				currentCoord.X = bf.Abs(currentShapesCoord.X)
			}
			if currentCoord.Y < bf.Abs(currentShapesCoord.Y) {
				currentCoord.Y = bf.Abs(currentShapesCoord.Y)
			}
		}
		for i := 0; i < 10; i++ {
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			for _, currentShapesCoord := range currentShapes.Body {
				x := currentCoord.X + currentShapesCoord.X
				y := currentCoord.Y + currentShapesCoord.Y
				termbox.SetCell(x, y, rune(' '), termbox.ColorGreen, termbox.ColorGreen)
			}
			termbox.Flush()
			currentCoord.Y++
			time.Sleep(time.Second)
		}
	}
}
