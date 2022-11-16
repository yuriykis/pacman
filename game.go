package main

import (
	"errors"
	"math/rand"
	"pacman/board"
	"pacman/character"
	"pacman/utils"
	"time"
)

type Game struct {
	characters []character.ICharacter
	items      []*Item
	engine     *Engine
}

func newGame() *Game {
	return &Game{engine: NewEngine()}
}

func (g *Game) createGame() {
	g.initPlayer()
	for i := 0; i < CharactersNumber; i++ {
		cType := rand.Intn(4) + 2
		pos := g.engine.board.FindFreePosition()
		char := character.NewCharacter(utils.CharacterType(cType))
		char.InitCharacter(pos)
		g.characters = append(g.characters, char)
	}
}

func (g *Game) initPlayer() {
	pos := g.engine.board.FindFreePosition()
	char := character.NewCharacter(utils.TPlayer)
	char.InitCharacter(pos)
	g.engine.player = char.(*character.Player)
}

func (g *Game) createBoard() {
	g.engine.board = board.NewBoard()
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
	go g.engine.movePlayer()
}

func (g *Game) startMoving(c character.ICharacter) {
	for {
		direction := c.Move()
		g.engine.moveCharacter(c, direction)
		time.Sleep(time.Duration(GameSpeed) * time.Millisecond)
	}
}
