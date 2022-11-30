package item

import (
	"pacman/board"
	"pacman/utils"

	"fyne.io/fyne/v2"
)

type IItem interface {
	ItemImage() *fyne.StaticResource
	ItemType() utils.ItemType
	InitItem(pos *board.Position)
	Position() *board.Position
}

type item struct {
	img   *fyne.StaticResource
	iType utils.ItemType
	pos   *board.Position
}

func (i *item) ItemImage() *fyne.StaticResource {
	return i.img
}

func (i *item) ItemType() utils.ItemType {
	return i.iType
}

func (i *item) InitItem(pos *board.Position) {
	i.setItemImage()
	i.SetItemPosition(pos)
}

func (i *item) setItemImage() {
	img, ok := utils.ResourceForItem(i.ItemType()).(*fyne.StaticResource)
	if ok {
		i.img = img
	}
}

func (i *item) SetItemPosition(pos *board.Position) {
	i.pos = pos
}

func (i *item) Position() *board.Position {
	return i.pos
}

func NewItem(iType utils.ItemType) IItem {
	switch iType {
	case utils.TCoin:
		return &Coin{item: item{iType: utils.TCoin}}
	}
	return nil
}
