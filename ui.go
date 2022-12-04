package main

import (
	"image/color"
	"pacman/board"
	"pacman/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type UserInterface struct {
	window fyne.Window
	grid   *fyne.Container
	game   *Game
}

func NewUI(window fyne.Window) *UserInterface {
	ui := &UserInterface{window: window}
	return ui
}

func (ui *UserInterface) createGrid() *fyne.Container {
	cells := make([]fyne.CanvasObject, 0)

	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			img := canvas.NewImageFromResource(nil)
			bg := canvas.NewRectangle(color.Gray{0x30})
			img.FillMode = canvas.ImageFillContain
			cell := container.NewMax(bg, img)
			cells = append(cells, cell)
			ui.game.engine.board.CreatePosition(x, y, cell)
		}
	}
	return container.New(&boardLayout{}, cells...)
}

func (ui *UserInterface) RefreshGrid() {
	for _, pos := range ui.game.engine.board.Positions() {
		cell := pos.Cell()
		img := cell.(*fyne.Container).Objects[1].(*canvas.Image)
		img.Resource = ui.positionImage(pos)
		if pos.PositionType() == board.Wall {
			img.Resource = utils.ResourceForWall()
		}
		img.Refresh()
	}
}

func (ui *UserInterface) positionImage(pos *board.Position) fyne.Resource {
	char, err := ui.game.engine.CharacterByPosition(pos)
	if err == nil {
		return char.CharacterImage()
	}
	return nil
}

func (ui *UserInterface) CreateUI() fyne.CanvasObject {
	ui.game.createBoard()
	ui.grid = ui.createGrid()
	ui.game.engine.board.SetPositionTypes()
	return ui.grid
}
