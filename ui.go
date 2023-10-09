package main

import (
	"errors"
	"image/color"

	"github.com/yuriykis/pacman/board"
	"github.com/yuriykis/pacman/character"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type UI struct {
	window        fyne.Window
	grid          *fyne.Container
	PositionDatas []*board.PositionData
}

func NewUI(window fyne.Window) *UI {
	ui := &UI{
		window:        window,
		PositionDatas: make([]*board.PositionData, BoardSize*BoardSize),
	}
	ui.createUI()
	return ui
}

func (ui *UI) createGrid() {
	cells := make([]fyne.CanvasObject, BoardSize*BoardSize)

	it := 0
	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			img := canvas.NewImageFromResource(nil)
			bg := canvas.NewRectangle(color.Gray{0x30})
			img.FillMode = canvas.ImageFillContain
			cell := container.NewMax(bg, img)
			cells[it] = cell
			ui.PositionDatas[it] = &board.PositionData{
				X:    x,
				Y:    y,
				Cell: cell,
			}
			it++
		}
	}
	ui.grid = container.New(&boardLayout{}, cells...)
}

func (ui *UI) RefreshGrid(poss []*board.Position, pImgs []fyne.Resource) error {
	for i, pos := range poss {
		cell := pos.GetCell()
		img, ok := cell.(*fyne.Container).Objects[1].(*canvas.Image)
		if !ok {
			return errors.New("error refreshing grid: could not cast cell to imag")
		}
		img.Resource = pImgs[i]
		if pos.PositionType() == board.Wall {
			img.Resource = character.ResourceForWall()
		}
		img.Refresh()
	}
	return nil
}

func (ui *UI) createUI() {
	ui.createGrid()
	ui.window.SetContent(ui.grid)
}
