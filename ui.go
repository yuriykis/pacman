package main

import (
	"image/color"
	"pacman/board"
	"pacman/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type userInterface struct {
	window fyne.Window
	grid   *fyne.Container
	game   *Game
}

func newUI(window fyne.Window) *userInterface {
	ui := &userInterface{window: window}
	return ui
}

func (ui *userInterface) createGrid() *fyne.Container {
	cells := make([]fyne.CanvasObject, 0)

	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			bg := canvas.NewRectangle(color.Gray{0x30})
			img := canvas.NewImageFromResource(nil)
			img.FillMode = canvas.ImageFillContain
			cell := container.NewMax(bg, img)
			cells = append(cells, cell)
			ui.game.createPosition(x, y, cell)
		}
	}
	return container.New(&boardLayout{}, cells...)
}

func (ui *userInterface) refreshGrid() {
	for _, pos := range ui.game.positions {
		cell := pos.Cell()
		img := cell.(*fyne.Container).Objects[1].(*canvas.Image)
		char, err := ui.game.CharacterByPosition(pos)
		if err == nil {
			img.Resource = char.CharacterImage()
		} else {
			img.Resource = nil
		}
		if pos.PositionType() == board.Wall {
			img.Resource = utils.ResourceForWall()
		}
		img.Refresh()
	}
}

func (ui *userInterface) createUI() fyne.CanvasObject {
	ui.grid = ui.createGrid()
	ui.game.createBoard()
	return ui.grid
}
