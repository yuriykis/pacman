package main

import (
	"pacman/character"
	"pacman/move"

	"github.com/pkg/term"
)

var (
	up     byte = 65
	down   byte = 66
	left   byte = 68
	right  byte = 67
	escape byte = 27
	keys        = map[byte]bool{
		up:    true,
		down:  true,
		left:  true,
		right: true,
	}
)

type Engine struct {
	player *character.Player
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) Run() {
	for {
		keyCode := getInput()
		switch keyCode {
		case up:
			e.player.SetDirection(move.Up)
		case down:
			e.player.SetDirection(move.Down)
		case left:
			e.player.SetDirection(move.Left)
		case right:
			e.player.SetDirection(move.Right)
		}
	}
}

func getInput() byte {
	t, _ := term.Open("/dev/tty")
	defer t.Close()

	err := term.RawMode(t)
	if err != nil {
		panic(err)
	}

	var read int
	readBytes := make([]byte, 3)
	read, err = t.Read(readBytes)

	if err != nil {
		panic(err)
	}

	t.Restore()
	if read == 3 {
		if _, ok := keys[readBytes[2]]; ok {
			return readBytes[2]
		}
	} else {
		return readBytes[0]
	}
	return 0
}
