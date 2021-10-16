package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element
	speed float32
}

func newKeyboardMover(container *element, speed float32) (*keyboardMover, error) {
	km := &keyboardMover{container: container, speed: speed}

	return km, nil
}


func (km *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (km *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		km.container.x -= int32(km.speed)
	}

	if keys[sdl.SCANCODE_RIGHT] == 1 {
		km.container.x += int32(km.speed)
	}

	return nil
}
