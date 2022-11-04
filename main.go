package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const Resolution int = 20
const CharactersNumber int = 50

func main() {
	application := app.New()
	window := application.NewWindow("Packman")
	ui := newUI(window)
	ui.game = newGame()
	window.SetContent(ui.createUI())
	window.Resize(fyne.NewSize(480, 480))

	// go func() {
	// 	for i := 0; i < Resolution; i++ {
	// 		for j := 0; j < Resolution; j++ {
	// 			//ui.refreshGrid(ui.grid, i, j)
	// 			//time.Sleep(time.Millisecond * 550)
	// 		}
	// 	}
	// }()
	ui.game.createGame()
	ui.refreshGrid(ui.grid)
	window.ShowAndRun()
}
