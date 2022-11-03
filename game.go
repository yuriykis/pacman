package main

import (
	"math/rand"

	"fyne.io/fyne/v2/canvas"
)

type Game struct {
	characters []*Character
	positions  []*Position
}

func newGame() *Game {
	return &Game{}
}

func (g *Game) findFreePosition() *Position {
	for _, p := range g.positions {
		if p.isFree {
			return p
		}
	}
	return nil
}
func (g *Game) createPosition(x int, y int) {
	pos := newPosition(x, y)
	g.positions = append(g.positions, pos)
}

func (g *Game) createGame() {
	g.initPlayer()

	for i := 0; i < CharactersNumber; i++ {
		cType := rand.Intn(3) + 1
		pos := g.findFreePosition()
		char := newCharacter(nil, CharacterType(cType))
		pos.assignCharacterToPosition(char)
		g.characters = append(g.characters, char)
	}
}

func (g *Game) initPlayer() {
	pos := g.findFreePosition()
	img := canvas.NewImageFromResource(resourceForCharacter(Player))
	img.FillMode = canvas.ImageFillContain
	char := newCharacter(img, CharacterType(0))
	pos.assignCharacterToPosition(char)
	g.characters = append(g.characters, char)
}

func (g *Game) PositionByCoords(x int, y int) *Position {
	for _, p := range g.positions {
		if p.X == x && p.Y == y {
			return p
		}
	}
	return nil
}
