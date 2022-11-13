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
