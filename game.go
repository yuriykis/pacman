package main

import (
	"math/rand"
	"pacman/board"
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
	//g.generateItems()
}

func (g *Game) generateCharacters() {
	for i := 0; i < CharactersNumber; i++ {
		cType := rand.Intn(4) + 2
		pos := g.engine.board.FindFreePosition()
		char := animated.NewAnimatedCharacter(types.CharacterType(cType))
		char.InitCharacter(pos)
		g.engine.characters = append(g.engine.characters, char)
	}
}

// func (g *Game) generateItems() {
// 	for i := 0; i < ItemsNumber; i++ {
// 		iType := rand.Intn(1) + 1
// 		pos := g.engine.board.FindFreePosition()
// 		it := item.NewItem(types.CharacterType(iType))
// 		it.InitItem(pos)
// 		g.engine.items = append(g.engine.items, it)
// 	}
// }

func (g *Game) initPlayer() {
	pos := g.engine.board.FindFreePosition()
	char := animated.NewAnimatedCharacter(types.TPlayer)
	char.InitCharacter(pos)
	g.engine.player = char.(*animated.Player)
}

func (g *Game) createBoard() {
	g.engine.board = board.NewBoard()
}

func (g *Game) StartGame() {
	for _, c := range g.engine.characters {
		go g.startMoving(c)
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
