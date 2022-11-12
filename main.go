package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const (
	BoardSize        int     = 20
	CharactersNumber int     = 20
	WindowResolution float32 = 480
)

func main() {
	application := app.New()
	window := application.NewWindow("Packman")
	window.Resize(fyne.NewSize(WindowResolution, WindowResolution))
	ui := newUI(window)
	ui.game = newGame()
	window.SetContent(ui.createUI())

	ui.game.createGame()

	go func() {
		for {
			ui.refreshGrid()
			ui.game.startGame()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	window.ShowAndRun()
}
