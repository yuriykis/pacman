package main

import (
	"math/rand"
	"pacman/board"
	"pacman/character"
	"pacman/character/animated"
	"pacman/character/types"
	"time"
)

type Game struct {
	engine *Engine
}

func NewGame() *Game {
	return &Game{engine: NewEngine()}
}

func (g *Game) CreateGame() {
	g.initPlayer()
	g.generateCharacters()
}

func (g *Game) generateCharacters() {
	for i := 0; i < CharactersNumber; i++ {
		cType := rand.Intn(5) + 2
		pos := g.engine.board.FindFreePosition()
		char := character.NewCharacter(types.CharacterType(cType))
		char.InitCharacter(pos)
		g.engine.characters = append(g.engine.characters, char)
	}
}

func (g *Game) initPlayer() {
	pos := g.engine.board.FindFreePosition()
	char := character.NewCharacter(types.TPlayer)
	char.InitCharacter(pos)
	g.engine.player = char.(*animated.Player)
}

func (g *Game) createBoard() {
	g.engine.board = board.NewBoard()
}

func (g *Game) StartGame() {
	for _, c := range g.engine.characters {
		char, ok := c.(animated.IAnimated)
		if ok {
			go g.startMoving(char)
		}
	}
	go g.engine.MovePlayer()
}

func (g *Game) startMoving(c animated.IAnimated) {
	for {
		direction := c.Move()
		g.engine.MoveCharacter(c, direction)
		time.Sleep(time.Duration(GameSpeed) * time.Millisecond)
	}
}
