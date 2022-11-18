package item

import (
	"pacman/board"
	"pacman/utils"

	"fyne.io/fyne/v2/canvas"
)

type IItem interface {
	ItemImage() *canvas.Image
	ItemType() utils.ItemType
}

type item struct {
	img   *canvas.Image
	iType utils.ItemType
	pos   *board.Position
}

func (i *item) ItemImage() *canvas.Image {
	return i.img
}

func (i *item) ItemType() utils.ItemType {
	return i.iType
}

func NewItem(iType utils.ItemType) IItem {
	switch iType {
	case utils.TCoin:
		return &Coin{item: item{iType: utils.TCoin}}
	}
	return nil
}
