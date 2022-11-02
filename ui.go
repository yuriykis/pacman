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
	grid := container.NewGridWithColumns(RESOLUTION)

	for y := 0; y < RESOLUTION; y++ {
		for x := 0; x < RESOLUTION; x++ {
			bg := canvas.NewRectangle(color.Gray{0x30})
			img := canvas.NewImageFromResource(nil)
			img.FillMode = canvas.ImageFillContain
			grid.Add(container.NewMax(bg, img))
			ui.game.createPosition(x, y)
		}
	}
	return grid
}

func (ui *userInterface) refreshGrid(grid *fyne.Container) {
	for _, cell := range grid.Objects {
		img := cell.(*fyne.Container).Objects[1].(*canvas.Image)
		img.Resource = resourceForCharacter(true)
		img.Refresh()
	}
}

func (ui *userInterface) createUI() fyne.CanvasObject {
	ui.grid = ui.createGrid()
	return ui.grid
}
