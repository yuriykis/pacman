package main

import (
	"pacman/board"
	"pacman/character"
	"pacman/move"
	"time"
)

type Engine struct {
	player *character.Player
	board  *board.Board
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) MovePlayer() {
	for {
		direction := e.player.Move()
		e.MoveCharacter(e.player, direction)
		time.Sleep(time.Duration(GameSpeed) * time.Millisecond)
	}
}

func (e *Engine) MoveCharacter(c character.ICharacter, direction move.Direction) {
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
		c.SetCharacterPosition(newPos)
		pos.SetPositionType(board.Space)
		pos.SetFree(true)
		newPos.SetFree(false)
	}
	newPos.Unlock()
}
