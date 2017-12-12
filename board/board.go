package board

type Board struct {
	Height     int
	Width      int
	GameHeight int
	GameWidth  int
}

func InitBoard() Board {
	return Board{30, 10, 30, 10}
}
