package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const RESOLUTION int = 20

func main() {
	application := app.New()
	window := application.NewWindow("Packman")

	ui := newUI(window)
	window.SetContent(ui.createUI())
	window.Resize(fyne.NewSize(480, 480))

	go func() {
		for i := 0; i < RESOLUTION; i++ {
			for j := 0; j < RESOLUTION; j++ {
				ui.refreshGrid(ui.grid, i, j)
				time.Sleep(time.Millisecond * 550)
			}
		}
	}()
	window.ShowAndRun()
}
