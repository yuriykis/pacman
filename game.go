package main

import (
	"errors"
	"fmt"
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

func (g *Game) CreateGame() error {
	err := g.initPlayer()
	if err != nil {
		return err
	}
	g.generateCharacters()
	return nil
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

func (g *Game) initPlayer() error {
	pos := g.engine.board.FindFreePosition()
	char := character.NewCharacter(types.TPlayer)
	char.InitCharacter(pos)
	player, ok := char.(*animated.Player)
	if !ok {
		return errors.New("could not create player")
	}
	g.engine.player = player
	return nil
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
	defer handleStartMovingPanic()
	for {
		direction := c.Move()
		g.engine.MoveCharacter(c, direction)
		time.Sleep(time.Duration(GameSpeed) * time.Millisecond)
	}
}

func handleStartMovingPanic() {
	if r := recover(); r != nil {
		fmt.Println("error in start moving func", r)
	}
}
