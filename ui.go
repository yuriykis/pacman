package main

import (
	"image/color"

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
	grid := container.NewGridWithColumns(Resolution)

	for y := 0; y < Resolution; y++ {
		for x := 0; x < Resolution; x++ {
			bg := canvas.NewRectangle(color.Gray{0x30})
			img := canvas.NewImageFromResource(nil)
			img.FillMode = canvas.ImageFillContain
			grid.Add(container.NewMax(bg, img))
			ui.game.createPosition(x, y)
		}
	}
	return grid
}

func (ui *userInterface) refreshGrid() {
	for _, cell := range ui.grid.Objects {
		x := cell.(*fyne.Container).Position().X
		y := cell.(*fyne.Container).Position().Y
		pos := ui.game.PositionByCoords(x, y)
		img := cell.(*fyne.Container).Objects[1].(*canvas.Image)
		char, err := pos.CharacterFromPosition()
		if err == nil {
			cType := char.cType
			img.Resource = resourceForCharacter(cType)
		}
		if pos.PositionType() == Wall {
			img.Resource = resourceForWall()
		}
		img.Refresh()
	}
}

func (ui *userInterface) createUI() fyne.CanvasObject {
	ui.grid = ui.createGrid()
	ui.game.createBoard()
	return ui.grid
}
