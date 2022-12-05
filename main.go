package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const (
	BoardSize        int     = 20
	CharactersNumber int     = 10
	WindowResolution float32 = 480
	GameSpeed        int     = 200
	ItemsNumber      int     = 10
)

func main() {
	application := app.New()
	window := application.NewWindow("Packman")
	window.Resize(fyne.NewSize(WindowResolution, WindowResolution))
	ui := NewUI(window)
	ui.game = NewGame()
	window.SetContent(ui.CreateUI())

	err := ui.game.CreateGame()
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			err := ui.RefreshGrid()
			if err != nil {
				panic(err)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
	ui.game.StartGame()

	window.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		ui.game.engine.player.MapKeyCodeToDirection(*k)
	})

	window.ShowAndRun()
}
