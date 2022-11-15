package main

import (
	"errors"
	"math/rand"
	"pacman/board"
	"pacman/character"
	"pacman/move"
	"pacman/utils"
	"time"

	"fyne.io/fyne/v2"
)

type Game struct {
	characters []character.ICharacter
	positions  []*board.Position
	items      []*Item
	board      *board.Board
	engine     *Engine
}

func newGame() *Game {
	return &Game{engine: NewEngine()}
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

func (g *Game) createPosition(x int, y int, cell fyne.CanvasObject) {
	pos := board.NewPosition(x, y, cell)
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
	g.engine.player = char.(*character.Player)
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
	if g.engine.player.Position() == pos {
		return g.engine.player, nil
	}
	for _, c := range g.characters {
		if c.Position().X == pos.X && c.Position().Y == pos.Y {
			return c, nil
		}
	}
	return nil, errors.New("no character on given position")
}

func (g *Game) startGame() {
	for _, c := range g.characters {
		go g.startMoving(c)
	}
	go g.movePlayer()
}

func (g *Game) startMoving(c character.ICharacter) {
	for {
		direction := c.Move()
		g.moveCharacter(c, direction)
		time.Sleep(time.Duration(GameSpeed) * time.Millisecond)
	}
}

func (g *Game) movePlayer() {
	for {
		direction := g.engine.player.Move()
		g.moveCharacter(g.engine.player, direction)
		time.Sleep(time.Duration(GameSpeed) * time.Millisecond)
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
	case move.NoDirection:
		return
	}
	newPos.Lock()
	if newPos.IsFree() {
		c.SetCharacterPosition(newPos)
		pos.SetPositionType(board.Space)
		pos.SetFree(true)
		newPos.SetFree(false)
	}
	newPos.Unlock()
}
