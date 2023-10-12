package main

import (
	"context"
	"fmt"
	"time"

	"github.com/yuriykis/pacman/character"

	"fyne.io/fyne/v2"
)

type Quit struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewQuit() *Quit {
	ctx, cancel := context.WithCancel(context.Background())
	return &Quit{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (q *Quit) Quit() {
	q.cancel()
}

func (q *Quit) Done() <-chan struct{} {
	return q.ctx.Done()
}

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
	q := NewQuit()
	go func(q *Quit) {
		for {
			select {
			case <-q.Done():
				g.ui.ShowDialog("Game over")
				return
			default:
				{
					g.Update()
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}(q)
	for _, c := range g.engine.characters {
		char, ok := c.(character.Mover)
		if ok {
			go g.startMoving(q, char)
		}
	}
	go g.engine.MovePlayer(q)
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

func (g *Game) startMoving(q *Quit, c character.Mover) {
	defer handleStartMovingPanic()
	for {
		select {
		case <-q.Done():
			return
		default:
			{
				direction := c.Move()
				g.engine.MoveCharacter(q, c.(character.BaseCharacter), direction)
				time.Sleep(time.Duration(GameSpeed) * time.Millisecond)
			}
		}
	}
}

func handleStartMovingPanic() {
	if r := recover(); r != nil {
		fmt.Println("error in start moving func", r)
	}
}
