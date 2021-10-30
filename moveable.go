package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type moveable struct {
	container *element
	maxVelocity float64
	drag float64
}

func newMoveable(container *element, maxVelocity float64, drag float64) (*moveable, error) {
	mv := &moveable{container: container, maxVelocity: maxVelocity, drag: drag}

	return mv, nil
}

func (*moveable) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mv *moveable) applyVelocity(velocity float64) float64 {
	if (velocity < 0) {
		velocity = math.Max(velocity, mv.maxVelocity * -1)
	} else {
		velocity = math.Min(velocity, mv.maxVelocity)
	}
	velocity *= 1 - gameDelta * mv.drag // fix the game delta locations all over.

	return velocity
}


func (mv *moveable) onUpdate() error {
	mv.container.velocity.x = mv.applyVelocity(mv.container.velocity.x)
	mv.container.velocity.y = mv.applyVelocity(mv.container.velocity.y)

	mv.container.position.x += gameDelta * mv.container.velocity.x
	mv.container.position.y += gameDelta * mv.container.velocity.y

	return nil
}
