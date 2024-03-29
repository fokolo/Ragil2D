package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element
	speed float64
}

func newKeyboardMover(container *element, speed float64) (*keyboardMover, error) {
	km := &keyboardMover{container: container, speed: speed}

	return km, nil
}


func (km *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (km *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		km.container.addVelocity(vector{x: -km.speed})
	}

	if keys[sdl.SCANCODE_RIGHT] == 1 {
		km.container.addVelocity(vector{x: km.speed})
	}

	if keys[sdl.SCANCODE_SPACE] == 1 {
		km.container.addVelocity(vector{y: -150})
	}

	return nil
}
