package main

import (
	"pacman/character"
)

type Engine struct {
	player *character.Player
}

func NewEngine() *Engine {
	return &Engine{}
}

// func (e *Engine) movePlayer() {
// 	for {
// 		direction := e.player.Move()
// 		e.moveCharacter(e.player, direction)
// 		time.Sleep(time.Duration(GameSpeed) * time.Millisecond)
// 	}
// }

// func (e *Engine) moveCharacter(c character.ICharacter, direction move.Direction) {
// 	pos := c.Position()
// 	var newPos *board.Position
// 	switch direction {
// 	case move.Up:
// 		newPos = g.findPosition(pos.X, pos.Y-1)
// 	case move.Down:
// 		newPos = g.findPosition(pos.X, pos.Y+1)
// 	case move.Left:
// 		newPos = g.findPosition(pos.X-1, pos.Y)
// 	case move.Right:
// 		newPos = g.findPosition(pos.X+1, pos.Y)
// 	case move.NoDirection:
// 		return
// 	}
// 	newPos.Lock()
// 	if newPos.IsFree() {
// 		c.SetCharacterPosition(newPos)
// 		pos.SetPositionType(board.Space)
// 		pos.SetFree(true)
// 		newPos.SetFree(false)
// 	}
// 	newPos.Unlock()
// }
