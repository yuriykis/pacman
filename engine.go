package main

import (
	"errors"
	"pacman/board"
	"pacman/character/animated"
	"pacman/character/base"
	"pacman/character/types"
	"pacman/move"
	"time"
)

type Engine struct {
	characters []base.ICharacter
	player     *animated.Player
	board      *board.Board
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) CharacterByPosition(pos *board.Position) (base.ICharacter, error) {
	if e.player.Position() == pos {
		return e.player, nil
	}
	for _, c := range e.characters {
		if c.Position().X == pos.X && c.Position().Y == pos.Y {
			return c, nil
		}
	}
	return nil, errors.New("no character on given position")
}

func (e *Engine) MovePlayer() {
	for {
		direction := e.player.Move()
		e.MoveCharacter(e.player, direction)
		time.Sleep(time.Duration(GameSpeed) * time.Millisecond)
	}
}

func (e *Engine) MoveCharacter(c animated.IAnimated, direction move.Direction) {
	pos := c.Position()
	var newPos *board.Position
	switch direction {
	case move.Up:
		newPos = e.board.FindPosition(pos.X, pos.Y-1)
	case move.Down:
		newPos = e.board.FindPosition(pos.X, pos.Y+1)
	case move.Left:
		newPos = e.board.FindPosition(pos.X-1, pos.Y)
	case move.Right:
		newPos = e.board.FindPosition(pos.X+1, pos.Y)
	case move.NoDirection:
		return
	}
	newPos.Lock()
	if newPos.IsFree() {
		e.moveToNewPosision(c, pos, newPos)
	} else {
		e.colision(c, newPos)
	}
	newPos.Unlock()
}

func (e *Engine) moveToNewPosision(c animated.IAnimated, pos *board.Position, newPos *board.Position) {
	c.SetCharacterPosition(newPos)
	pos.SetPositionType(board.Space)
	pos.SetFree(true)
	newPos.SetFree(false)
}

func (e *Engine) colision(myChar animated.IAnimated, newPos *board.Position) {
	if newPos.PositionType() == board.Wall {
		return
	}
	if myChar.CharacterType() == types.TPlayer {
		// colisionCharacter, err := e.CharacterByPosition(newPos)
		// if err != nil {
		// 	return
		// }
		// if colisionCharacter.AnimatedType() == utils.TCoin {
		// 	//	e.player.AddScore(10)
		// 	//			e.removeCharacter(colisionCharacter)
		// 	return
		// }
		return
	}
}
