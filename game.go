package main

import (
	"errors"
	"math/rand"
	"pacman/board"
	"pacman/character"
	"pacman/move"
	"pacman/utils"
	"time"
)

type Game struct {
	characters []character.ICharacter
	items      []*Item
	board      *board.Board
	engine     *Engine
}

func newGame() *Game {
	return &Game{engine: NewEngine()}
}

func (g *Game) createGame() {
	g.initPlayer()
	for i := 0; i < CharactersNumber; i++ {
		cType := rand.Intn(4) + 2
		pos := g.board.FindFreePosition()
		char := character.NewCharacter(utils.CharacterType(cType))
		char.InitCharacter(pos)
		g.characters = append(g.characters, char)
	}
}

func (g *Game) initPlayer() {
	pos := g.board.FindFreePosition()
	char := character.NewCharacter(utils.TPlayer)
	char.InitCharacter(pos)
	g.engine.player = char.(*character.Player)
}

func (g *Game) createBoard() {
	g.board = board.NewBoard()
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
		newPos = g.board.FindPosition(pos.X, pos.Y-1)
	case move.Down:
		newPos = g.board.FindPosition(pos.X, pos.Y+1)
	case move.Left:
		newPos = g.board.FindPosition(pos.X-1, pos.Y)
	case move.Right:
		newPos = g.board.FindPosition(pos.X+1, pos.Y)
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
