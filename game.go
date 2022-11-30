package main

import (
	"errors"
	"math/rand"
	"pacman/board"
	"pacman/character"
	"pacman/item"
	"pacman/utils"
	"time"
)

type Game struct {
	characters []character.ICharacter
	items      []item.IItem
	engine     *Engine
}

func newGame() *Game {
	return &Game{engine: NewEngine()}
}

func (g *Game) createGame() {
	g.initPlayer()
	g.generateCharacters()
	g.generateItems()
}

func (g *Game) generateCharacters() {
	for i := 0; i < CharactersNumber; i++ {
		cType := rand.Intn(4) + 2
		pos := g.engine.board.FindFreePosition()
		char := character.NewCharacter(utils.CharacterType(cType))
		char.InitCharacter(pos)
		g.characters = append(g.characters, char)
	}
}

func (g *Game) generateItems() {
	for i := 0; i < ItemsNumber; i++ {
		iType := rand.Intn(1) + 1
		pos := g.engine.board.FindFreePosition()
		it := item.NewItem(utils.ItemType(iType))
		it.InitItem(pos)
		g.items = append(g.items, it)
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

func (g *Game) ItemByPosition(pos *board.Position) (item.IItem, error) {
	for _, i := range g.items {
		if i.Position().X == pos.X && i.Position().Y == pos.Y {
			return i, nil
		}
	}
	return nil, errors.New("no item on given position")
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
