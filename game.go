package main

import (
	"fmt"
	"time"

	"github.com/yuriykis/pacman/character"

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
	g.createBoardPositions()
	return g
}

func (g *Game) CreateGame() error {
	err := g.engine.initPlayer()
	if err != nil {
		return err
	}
	g.engine.generateCharacters()
	return nil
}

func (g *Game) Start() error {
	err := g.CreateGame()
	if err != nil {
		return err
	}

	go func() {
		for {
			g.Update()
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

func (g *Game) Update() {
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

func (g *Game) createBoardPositions() {
	for _, pd := range g.ui.PositionDatas {
		g.engine.board.CreatePosition(pd.X, pd.Y, pd.Cell)
	}
	g.engine.board.SetPositionTypes()
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
