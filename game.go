package main

import (
	"errors"
	"math/rand"
	"pacman/board"
	"pacman/character"
	"pacman/move"
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
		if p.IsFree() {
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

func (g *Game) PositionByBoardCoords(x float32, y float32) *board.Position {
	var coef int = int(WindowResolution) / BoardSize
	for _, p := range g.positions {
		if (int(x)/coef) == p.X && (int(y)/coef) == p.Y {
			return p
		}
	}
	return nil
}

func (g *Game) findPosition(x int, y int) *board.Position {
	for _, p := range g.positions {
		if p.X == x && p.Y == y {
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

func (g *Game) startGame() {
	for _, c := range g.characters {
		direction := c.Move()
		g.moveCharacter(c, direction)
	}
}

func (g *Game) moveCharacter(c character.ICharacter, direction move.Direction) {
	pos := c.Position()
	var newPos *board.Position
	switch direction {
	case move.Up:
		newPos = g.findPosition(pos.X, pos.Y-1)
	case move.Down:
		newPos = g.findPosition(pos.X, pos.Y+1)
	case move.Left:
		newPos = g.findPosition(pos.X-1, pos.Y)
	case move.Right:
		newPos = g.findPosition(pos.X+1, pos.Y)
	}
	if newPos.IsFree() {
		c.SetCharacterPosition(newPos)
		pos.SetPositionType(board.Space)
		pos.SetFree(true)
		newPos.SetFree(false)
	}
}
