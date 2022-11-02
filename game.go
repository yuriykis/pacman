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

func (g *Game) isFreePosition(pos Position) bool {
	for _, c := range g.characters {
		if c.pos.X == pos.X && c.pos.Y == pos.Y {
			return false
		}
	}
	return true
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
	pos := newPosition(x, y, 0, true)
	g.positions = append(g.positions, pos)
}

func (g *Game) createGame() {
	g.initPlayer()

	pos := g.findFreePosition()
	cType := rand.Intn(3) + 1
	for i := 0; i < CHARACTERS_NUMBER; i++ {
		char := newCharacter(pos, nil, CharacterType(cType))
		g.characters = append(g.characters, char)
	}
}

func (g *Game) initPlayer() {
	pos := g.findFreePosition()
	img := canvas.NewImageFromResource(resourceForCharacter(true))
	img.FillMode = canvas.ImageFillContain
	char := newCharacter(pos, img, CharacterType(0))
	g.characters = append(g.characters, char)
}
