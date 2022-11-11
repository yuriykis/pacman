package main

import (
	"errors"
	"math/rand"
	"pacman/board"
	"pacman/character"
	"pacman/utils"
)

type Game struct {
	characters []character.ICharacter
	positions  []*board.Position
	items      []*Item
	board      *board.Board
}

func newGame() *Game {
	return &Game{}
}

func (g *Game) findFreePosition() *board.Position {
	var freePositions []*board.Position
	for _, p := range g.positions {
		if p.IsFree {
			freePositions = append(freePositions, p)
		}
	}
	if len(freePositions) != 0 {
		return freePositions[rand.Intn(len(freePositions))]
	}
	return nil
}

func (g *Game) createPosition(x int, y int) {
	pos := board.NewPosition(x, y)
	g.positions = append(g.positions, pos)
}

func (g *Game) createGame() {
	g.initPlayer()
	for i := 0; i < CharactersNumber; i++ {
		cType := rand.Intn(4) + 2
		pos := g.findFreePosition()
		char := character.NewCharacter(utils.CharacterType(cType))
		char.InitCharacter(pos)
		g.characters = append(g.characters, char)
	}
}

func (g *Game) initPlayer() {
	pos := g.findFreePosition()
	char := character.NewCharacter(utils.TPlayer)
	char.InitCharacter(pos)
	g.characters = append(g.characters, char)
}

func (g *Game) PositionByCoords(x float32, y float32) *board.Position {
	for _, p := range g.positions {
		if (int(x)/5) == p.X && (int(y)/5) == p.Y {
			return p
		}
	}
	return nil
}

func (g *Game) createBoard() {
	g.board = board.NewBoard()
	for _, p := range g.positions {
		if g.board.BoardPositionTypeByPosition(*p) == board.WallChar {
			p.SetPositionType(board.Wall)
			continue
		}
		if g.board.BoardPositionTypeByPosition(*p) == board.SpaceChar {
			p.SetPositionType(board.Space)
			continue
		}
	}
}

func (g *Game) CharacterByPosition(pos *board.Position) (character.ICharacter, error) {
	for _, c := range g.characters {
		if c.Position().X == pos.X && c.Position().Y == pos.Y {
			return c, nil
		}
	}
	return nil, errors.New("no character on given position")
}
