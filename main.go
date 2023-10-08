package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	log "github.com/sirupsen/logrus"
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
	window := application.NewWindow("Pacman")
	window.Resize(fyne.NewSize(WindowResolution, WindowResolution))
	game := NewGame(window)

	if err := game.Start(); err != nil {
		log.Fatal(err)
	}
}
