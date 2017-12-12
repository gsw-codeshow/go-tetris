package main

import (
	"go-tetris/block"
	"go-tetris/board"

	"fmt"
	"math/rand"

	"time"

	"github.com/nsf/termbox-go"
)

type Coord struct {
	X, Y int
}

type Context struct {
	CurrentBShapesPoint Coord
	CurrentShapes       block.Block
	Context             [][]int
	Event               chan termbox.Event
	Block               *block.BlockInterface
	Board               board.Board
	Quit                bool
}

func ListenEvent(event chan termbox.Event) {
	for {
		event <- termbox.PollEvent()
	}
	return
}

func (ct *Context) InitContextBoard() {
	for i := 0; i < ct.Board.Height; i++ {
		var boardTmp []int
		for j := 0; j < ct.Board.Width; j++ {
			boardTmp = append(boardTmp, 0)
		}
		ct.Context = append(ct.Context, boardTmp)
	}
}

func (ct *Context) RandomShapes() int {
	blockIndex := len(ct.Block.Entry)
	return rand.Int() % blockIndex
}

func (ct *Context) Draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for x := 0; x < ct.Board.Height; x++ {
		for y := 0; y < ct.Board.Width; y++ {
			if 0 != ct.Context[x][y] {
				termbox.SetCell(x, y, rune('#'), termbox.ColorGreen, termbox.ColorDefault)
			}
		}
	}
	for _, shapesCoord := range ct.CurrentShapes.Body {
		x := shapesCoord.X + ct.CurrentBShapesPoint.X
		y := shapesCoord.Y + ct.CurrentBShapesPoint.Y
		termbox.SetCell(x, y, rune('#'), termbox.ColorGreen, termbox.ColorDefault)
	}
	termbox.Flush()
}

func (ct *Context) InitShapesAndShapes() {
	ct.CurrentShapes = ct.Block.Entry[ct.RandomShapes()]
	ct.CurrentBShapesPoint = Coord{ct.Board.Width / 2, 0}
	maxInitPoint := Coord{0, 0}
	for _, blockOffset := range ct.CurrentShapes.Body {
		if maxInitPoint.X < ct.Block.Abs(blockOffset.X) {
			maxInitPoint.X = ct.Block.Abs(blockOffset.X)
		}
		if maxInitPoint.Y < ct.Block.Abs(blockOffset.Y) {
			maxInitPoint.Y = ct.Block.Abs(blockOffset.Y)
		}
	}
	ct.CurrentBShapesPoint.X += maxInitPoint.X
	ct.CurrentBShapesPoint.Y += maxInitPoint.Y
}

func (ct *Context) HasVaild() bool {
	for _, blockOffset := range ct.CurrentShapes.Body {
		x := ct.CurrentBShapesPoint.X + blockOffset.X
		y := ct.CurrentBShapesPoint.Y + blockOffset.Y
		if x > ct.Board.Width {
			return false
		}
		if y > ct.Board.Height {
			return false
		}
		if 0 != ct.Context[x][y] {
			return false
		}
	}
	return true
}

func main() {
	termboxErr := termbox.Init()
	if nil != termboxErr {
		fmt.Println(termboxErr)
	}
	defer termbox.Close()
	context := new(Context)
	context.Event = make(chan termbox.Event)

	go ListenEvent(context.Event)
	context.Block = block.InitBlockInterface()
	context.Board = board.InitBoard()
	context.InitContextBoard()
	context.InitShapesAndShapes()
	context.Quit = true
	for context.Quit {
		context.CurrentBShapesPoint.Y++
		context.Draw()
		time.Sleep(time.Second)
		if !context.HasVaild() {
		}
	}
}
