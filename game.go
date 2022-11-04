package main

import (
	"math/rand"

	"fyne.io/fyne/v2/canvas"
)

type Game struct {
	characters []*Character
	positions  []*Position
	board      *Board
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
	g.createBoard()
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

func (g *Game) PositionByCoords(x float32, y float32) *Position {
	for _, p := range g.positions {
		if (int(x)/5) == p.X && (int(y)/5) == p.Y {
			return p
		}
	}
	return nil
}

func (g *Game) createBoard() {
	g.board = NewBoard()
	for _, p := range g.positions {
		if g.board.BoardPositionTypeByPosition(*p) == WallChar {
			p.SetPositionType(Wall)
			continue
		}
		if g.board.BoardPositionTypeByPosition(*p) == SpaceChar {
			p.SetPositionType(Space)
			continue
		}
	}
}
