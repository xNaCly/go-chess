// field implements the logic for moving pieces around the chess board
package field

type Board struct {
	Field [8][8]Figure
}

// Init places the default pieces on the board
func (b *Board) Init() {
	panic("TODO")
}
