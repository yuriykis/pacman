package main

import (
	"errors"
	"fmt"
	"math/rand"
	"pacman/board"
	"pacman/character"
	"time"

	"fyne.io/fyne/v2"
)

type Game struct {
	engine *Engine
	ui     *UI
}

func NewGame(window fyne.Window) *Game {
	g := &Game{
		engine: NewEngine(),
	}
	g.ui = NewUI(window)
	g.createBoard()
	g.CreateBoardPositions()
	g.engine.board.SetPositionTypes()
	return g
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
		var (
			char  character.BaseCharacter
			cType = rand.Intn(5) + 2
			pos   = g.engine.board.FindFreePosition()
		)
		if cType < 5 {
			char = character.NewCharacter(character.CharacterType{
				MoverType: character.MoverType(cType),
			})
		} else {
			char = character.NewCharacter(character.CharacterType{
				CollectibleType: character.CoinType,
			})
		}
		char.InitCharacter(pos)
		g.engine.characters = append(g.engine.characters, char)
	}
}

func (g *Game) initPlayer() error {
	pos := g.engine.board.FindFreePosition()
	char := character.NewCharacter(character.CharacterType{
		MoverType:       character.PlayerType,
		CollectibleType: character.NoneCollectibleType,
	})
	char.InitCharacter(pos)
	player, ok := char.(*character.Player)
	if !ok {
		return errors.New("could not create player")
	}
	g.engine.player = player
	return nil
}

func (g *Game) createBoard() {
	g.engine.board = board.NewBoard()
}

func (g *Game) CreateBoardPositions() {
	for _, pd := range g.ui.PositionDatas {
		g.engine.board.CreatePosition(pd.X, pd.Y, pd.Cell)
	}
}

func (g *Game) Refresh() {
	poss := g.engine.board.Positions()
	pImgs := make([]fyne.Resource, len(poss))
	for i, pos := range poss {
		c, err := g.engine.CharacterByPosition(pos)
		if err == nil {
			pImgs[i] = c.CharacterImage()
			continue
		}
	}
	err := g.ui.RefreshGrid(poss, pImgs)
	if err != nil {
		panic(err)
	}
}

func (g *Game) Start() error {
	err := g.CreateGame()
	if err != nil {
		return err
	}

	go func() {
		for {
			g.Refresh()
			time.Sleep(100 * time.Millisecond)
		}

	}()

	for _, c := range g.engine.characters {
		char, ok := c.(character.Mover)
		if ok {
			go g.startMoving(char)
		}
	}
	go g.engine.MovePlayer()
	g.ui.window.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		g.engine.player.MapKeyCodeToDirection(*k)
	})
	g.ui.window.ShowAndRun()
	return nil
}

func (g *Game) startMoving(c character.Mover) {
	defer handleStartMovingPanic()
	for {
		direction := c.Move()
		g.engine.MoveCharacter(c.(character.BaseCharacter), direction)
		time.Sleep(time.Duration(GameSpeed) * time.Millisecond)
	}
}

func handleStartMovingPanic() {
	if r := recover(); r != nil {
		fmt.Println("error in start moving func", r)
	}
}
