package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type boardLayout struct{}

func (layout *boardLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	smallEdge := size.Width
	if size.Height < size.Width {
		smallEdge = size.Height
	}
	leftInset := (size.Width - smallEdge) / 2
	cellEdge := smallEdge / float32(BoardSize)
	cellSize := fyne.NewSize(cellEdge, cellEdge)
	i := 0
	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			objects[i].Resize(cellSize)
			objects[i].Move(fyne.NewPos(leftInset+float32(x)*cellEdge, float32(y)*cellEdge))
			i++
		}
	}
}

func (layout *boardLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	edge := theme.IconInlineSize() * float32(BoardSize)
	return fyne.NewSize(edge, edge)
}
