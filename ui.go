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
			var isUser bool
			if y == 0 && x == 0 {
				isUser = true
			} else {
				isUser = false
			}
			img := canvas.NewImageFromResource(resourceForCharacter(isUser))
			img.FillMode = canvas.ImageFillContain
			grid.Add(container.NewMax(bg, img))
		}
	}
	return grid
}

func (ui *userInterface) refreshGrid(grid *fyne.Container, i int, j int) {
	for it, cell := range grid.Objects {
		var isUser bool
		if it == ((i * RESOLUTION) + j) {
			isUser = true
		} else {
			isUser = false
		}
		img := cell.(*fyne.Container).Objects[1].(*canvas.Image)
		img.Resource = resourceForCharacter(isUser)
		img.Refresh()
	}
}

func (ui *userInterface) createUI() fyne.CanvasObject {
	ui.grid = ui.createGrid()
	return ui.grid
}
