package block

import (
	"reflect"
)

type BlockOffset struct {
	X, Y int
}

type Block struct {
	Body []BlockOffset
}

type BlockInterface struct {
	Entry []Block
}

func (bf *BlockInterface) Add(b Block) {
	bf.Entry = append(bf.Entry, b)
}

func (bf *BlockInterface) Delete(b Block) {
	var entry []Block
	for _, bkData := range bf.Entry {
		if reflect.DeepEqual(b.Body, bkData.Body) {
			continue
		}
		entry = append(entry, bkData)
	}
	bf.Entry = entry
}

func InitBlockInterface() *BlockInterface {
	bf := new(BlockInterface)
	TShapes := Block{Body: []BlockOffset{{-1, 1}, {0, 1}, {1, 1}, {0, 0}}}
	LShapes := Block{Body: []BlockOffset{{0, 0}, {0, 1}, {1, 0}, {2, 0}}}
	LineShapes := Block{Body: []BlockOffset{{-1, 0}, {0, 0}, {1, 0}, {2, 0}}}
	RELShapes := Block{Body: []BlockOffset{{1, -1}, {-1, 0}, {0, 0}, {1, 0}}}
	RessShapes := Block{Body: []BlockOffset{{0, -1}, {1, -1}, {-1, 0}, {0, 0}}}
	LessShapes := Block{Body: []BlockOffset{{-1, -1}, {0, -1}, {0, 0}, {1, 0}}}
	SquareShapes := Block{Body: []BlockOffset{{0, -1}, {1, -1}, {0, 0}, {1, 0}}}

	bf.Entry = append(bf.Entry, TShapes, LShapes, LineShapes, RELShapes, RessShapes, LessShapes, SquareShapes)
	return bf
}

func (bf *BlockInterface) Abs(x int) int {
	if x > 0 {
		return x
	}
	if x == 0 {
		return 0
	}
	return -x
}

func (bf *BlockInterface) Rotate(currentShapes Block) Block {
	var rotateShapes Block
	for _, currentBodyItem := range currentShapes.Body {
		x := -1 * currentBodyItem.Y
		y := currentBodyItem.X
		rotateShapes.Body = append(rotateShapes.Body, BlockOffset{x, y})
	}
	return rotateShapes
}
