package main

import (
	"github.com/veandco/go-sdl2/sdl"
)


func newPlayer(renderer *sdl.Renderer) *element {
	collisionCircles := []circle{{vector{100,500}, 30}, {vector{200,400}, 30}, {vector{500,300}, 30}, {vector{600,100}, 30}, }
	player := &element{
		position: vector{x: 350, y: 250},
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

	mv, err := newMoveable(player, 10.0, 0.05)
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
