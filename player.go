package main

import (
	"github.com/veandco/go-sdl2/sdl"
)


func newPlayer(renderer *sdl.Renderer) *element {
	player := &element{
		x: 700,
		y: 500,
		scale: 5,
		active: true,
	}

	sr, err := newSpriteRenderer(renderer, "sprites/TempleRun/Idle__000.png", player)
	if err != nil {
		panic("Error creating sprite renderer")
	}

	player.addComponent(sr)

	return player
}
