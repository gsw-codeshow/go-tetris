package block

import (
	_ "fmt"
	"reflect"
	"testing"
)

func Test_RotateTest(t *testing.T) {
	bf := InitBlockInterface()
	currentBlock := bf.Rotate(bf.Entry[1])
	if !reflect.DeepEqual(currentBlock.Body, []BlockOffset{{0, 0}, {-1, 0}, {0, 1}, {0, 2}}) {
		t.Fail()
	}
}
