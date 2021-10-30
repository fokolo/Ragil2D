package main

import (
	"github.com/veandco/go-sdl2/sdl"
)


func newPlayer(renderer *sdl.Renderer) *element {
	collisionCircles := []circle{{vector{0,screenHeight}, 30}, {vector{screenWidth,screenHeight}, 30}}
	player := &element{
		position: vector{x: 700, y: 550},
		radius: 40,
		scale: 5,
		active: true,
		collisionCircles: collisionCircles,
	}

	sr, err := newSpriteRenderer(player, renderer, "sprites/TempleRun/Idle__000.png")
	if err != nil {
		panic("Error creating sprite renderer")
	}

	km, err := newKeyboardMover(player, 0.5)
	if err != nil {
		panic("Error creating keyboard mover")
	}

	mv, err := newMoveable(player, 5.0, 0.05)
	if err != nil {
		panic("Error creating movable")
	}


	collisions, err := newCollision(player)
	if err != nil {
		panic("Error creating new collision")
	}

	player.addComponent(sr)
	player.addComponent(km)
	player.addComponent(mv)
	player.addComponent(collisions)

	return player
}
