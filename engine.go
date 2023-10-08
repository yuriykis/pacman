package main

import (
	"errors"
	"fmt"
	"math/rand"
	"pacman/board"
	"pacman/character"
	"pacman/move"
	"time"
)

type Engine struct {
	characters []character.BaseCharacter
	player     *character.Player
	board      *board.Board
}

func NewEngine() *Engine {
	return &Engine{
		board: board.NewBoard(),
	}
}

func (e *Engine) initPlayer() error {
	pos := e.board.FindFreePosition()
	char := character.NewCharacter(character.CharacterType{
		MoverType:       character.PlayerType,
		CollectibleType: character.NoneCollectibleType,
	})
	char.InitCharacter(pos)
	player, ok := char.(*character.Player)
	if !ok {
		return errors.New("could not create player")
	}
	e.player = player
	return nil
}

func (e *Engine) generateCharacters() {
	for i := 0; i < CharactersNumber; i++ {
		var (
			char  character.BaseCharacter
			cType = rand.Intn(5) + 2
			pos   = e.board.FindFreePosition()
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
		e.characters = append(e.characters, char)
	}
}

func (e *Engine) CharacterByPosition(
	pos *board.Position,
) (character.BaseCharacter, error) {
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
	defer handleMovePlayerPanic()
	for {
		direction := e.player.Move()
		e.MoveCharacter(e.player, direction)
		time.Sleep(time.Duration(GameSpeed) * time.Millisecond)
	}
}

func handleMovePlayerPanic() {
	if r := recover(); r != nil {
		fmt.Println("error in move player func", r)
	}
}

func (e *Engine) MoveCharacter(c character.BaseCharacter, direction move.Direction) {
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
		e.colision(c, pos, newPos)
	}
	newPos.Unlock()
}

func (e *Engine) moveToNewPosision(
	c character.BaseCharacter,
	pos *board.Position,
	newPos *board.Position,
) {
	c.SetCharacterPosition(newPos)
	pos.SetPositionType(board.Space)
	pos.SetFree(true)
	newPos.SetFree(false)
}

func (e *Engine) removeCharacter(c character.BaseCharacter) {
	for i, ch := range e.characters {
		if ch == c {
			e.characters = append(e.characters[:i], e.characters[i+1:]...)
			return
		}
	}
}

func (e *Engine) colision(
	myChar character.BaseCharacter,
	pos *board.Position,
	newPos *board.Position,
) {
	if newPos.PositionType() == board.Wall {
		return
	}
	if myChar.Type().MoverType != character.NoneMoverType {
		colisionCharacter, err := e.CharacterByPosition(newPos)
		if err != nil {
			return
		}
		if colisionCharacter.Type().CollectibleType != character.NoneCollectibleType {
			//	e.player.AddScore(10)
			e.moveToNewPosision(myChar, pos, newPos)
			e.removeCharacter(colisionCharacter)
			return
		}
		return
	}
}
