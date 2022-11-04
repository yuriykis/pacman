package main

import (
	"fyne.io/fyne/v2/canvas"
)

type Item struct {
	img   *canvas.Image
	iType ItemType
}

type ItemType int8

const (
	Box ItemType = iota
)

func newItem(img *canvas.Image, iType ItemType) *Item {
	return &Item{img: img, iType: iType}
}
